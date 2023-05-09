package postGresql

import (
	"btl/infra/model"
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

func (tk *collection) CanCelTicket(ctx context.Context, phone_number string, booking_id string) (bool, error) {

	booking := model.Booking{}
	result := tk.db.Model(&booking).Where("phone_number = ? and booking_id = ?", phone_number, booking_id).Updates(
		model.Booking{Status: "ticket canceled"})
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, errors.New("flight not found")
	}
	_, err := tk.TriggerAugment(ctx)
	if err != nil {
		// handle error
		return false, err
	}
	return true, nil
}
func (tk *collection) TriggerAugment(ctx context.Context) (bool, error) {
	// Increment the remaining_seats column in the flights table by 1
	result := tk.db.Model(&model.Flight{}).Update("remaining_seats", gorm.Expr("remaining_seats + ?", 1))
	if result.Error != nil {
		return false, nil
	}
	// log a message indicating that the trigger has been executed
	log.Println("Trigger executed successfully")
	return true, nil
}

func (tk *collection) GetStatusTicket(ctx context.Context, phone_number string, booking_id string) (*model.Booking, error) {
	booking := &model.Booking{}

	result := tk.db.Where("phone_number = ? and booking_id = ?", phone_number, booking_id).First(booking)
	if result.Error != nil {
		// If there was an error querying the database, return the error
		return nil, result.Error
	}

	return booking, nil
}

// dau vao se la ten chuyen bay , id chuyen bay
func (tk *collection) CreateTicket(ctx context.Context, ticket *model.BookingRequest) (bool, error) {

	flight, err := tk.GetByFlightById(ctx, ticket.FlightID)
	if err != nil {
		return false, err
	}
	user, err := tk.FindByNumber(ctx, ticket.PhoneNumber)
	if err != nil {
		return false, err
	}
	registerTime := time.Now()
	strTime := registerTime.Format("2006-01-02 15:04:05") // chuyển đổi thành chuỗi theo định dạng mong muốn

	// Tạo đối tượng Booking mới với thông tin từ BookingRequest và thông tin từ chuyến bay và người dùng được truy vấn
	var booking = model.Booking{
		BookingID:     uuid.NewString(),
		UserID:        user.UserID,
		FlightID:      ticket.FlightID,
		NumberOfSeats: ticket.NumberOfSeats,
		Amount:        flight.Fare * float64(ticket.NumberOfSeats),
		// Thông tin người đặt vé
		UserName:    user.UserName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		NumberCMND:  user.NumberCMND,
		Nationality: user.Nationality,
		Language:    user.Language,
		// Thông tin chuyến bay
		NameAirline:     flight.NameAirline,
		Destination:     flight.Destination,
		Departure:       flight.Departure,
		DestinationTime: flight.DestinationTime,
		DepartureTime:   flight.DepartureTime,
		TicketType:      ticket.TicketType,
		Fare:            flight.Fare,
		Status:          "Booked",
		NameFlight:      flight.Name_flight,
		RegisterTime:    strTime,
	} // Thêm đối tượng Booking vào cơ sở dữ liệu
	tk.db.Create(booking)

	// Cập nhật thông tin số lượng chỗ trống trên chuyến bay
	flight.RemainingSeats -= 1
	err = tk.UpdateFlightById(ctx, flight)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (tk *collection) TriggerReduce(ctx context.Context) (bool, error) {
	result := tk.db.Model(&model.Flight{}).Where("remaining_seats > 0").Update("remaining_seats", gorm.Expr("remaining_seats - ?", 1))
	if result.Error != nil {
		return false, nil
	}
	log.Println("TriggerReduce executed successfully")
	return true, nil
}

func (tk *collection) UpdateFlightById(ctx context.Context, flight *model.Flight) error {
	err := tk.db.Model(&model.Flight{}).Where("flight_id = ?", flight.FlightID).Updates(flight).Error
	if err != nil {
		return err
	}
	return nil
}

func (tk *collection) GetAllTicket(ctx context.Context) ([]*model.Booking, error) {
	var bookings []*model.Booking
	err := tk.db.Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
func (tk *collection) GetTicketByPhoneNumber(ctx context.Context, phone_number string) (*model.Booking, error) {
	var booking model.Booking
	err := tk.db.Where("phone_number = ?", phone_number).First(&booking).Error
	if err != nil {
		return nil, err
	}
	return &booking, nil
}
