package postGresql

import (
	"btl/infrastructure/model"
	"context"
	"errors"

	"github.com/google/uuid"
)

func (p *collection) CreateFlight(ctx context.Context, flight *model.Flight) (bool, error) {
	flight.FlightID = uuid.New().String()
	result := p.db.Create(flight)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// //////////
func (p *collection) UpdateFlight(ctx context.Context, flightID string, nameFlight string, updateFlight *model.Flight) (bool, error) {
	flight := model.Flight{}
	result := p.db.Model(&flight).Where("flight_id = ? AND name_flight = ?", flightID, nameFlight).Updates(model.Flight{
		NameAirline:     updateFlight.NameAirline,
		Destination:     updateFlight.Destination,
		Departure:       updateFlight.Departure,
		DestinationTime: updateFlight.DestinationTime,
		DepartureTime:   updateFlight.DepartureTime,
		RemainingSeats:  updateFlight.RemainingSeats,
		TicketType:      updateFlight.TicketType,
		Fare:            updateFlight.Fare,
		Name_flight:     updateFlight.Name_flight,
	})
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, errors.New("flight not found")
	}
	return true, nil
}
func (p *collection) DeleteFlight(ctx context.Context, flightID string, nameFlight string) (bool, error) {
	result := p.db.Delete(&model.Flight{}, "flight_id = ? AND name_flight = ?", flightID, nameFlight)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, errors.New("flight not found")
	}
	return true, nil
}

// ///////////////
//func (p *collection) FindAllFlight(ctx context.Context) ([]*model.Flight, error) {
//	var flights []*model.Flight
//	// Thực hiện câu truy vấn để lấy thông tin tất cả các chuyến bay trong bảng `flights`
//	result := p.db.Find(&flights)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return flights, nil
//}

func (p *collection) FindBFlightByForm(ctx context.Context, flight model.FlightByForm) ([]*model.Flight, error) {
	data := make([]*model.Flight, 0)
	if result := p.db.Where(model.Flight{
		FlightID:        flight.FlightID,
		NameAirline:     flight.NameAirline,
		Destination:     flight.Destination,
		Departure:       flight.Departure,
		DestinationTime: flight.DestinationTime,
		DepartureTime:   flight.DepartureTime,
		RemainingSeats:  flight.RemainingSeats,
		TicketType:      flight.TicketType,
		Fare:            flight.Fare,
		Status:          flight.Status,
		Name_flight:     flight.Name_flight,
	}).Find(&data).Error; result != nil {
		return nil, result
	}
	return data, nil
}
func (p *collection) GetByFlightById(ctx context.Context, flight_id string) (*model.Flight, error) {
	var flight model.Flight
	result := p.db.Where("flight_id  = ?", flight_id).First(&flight)
	if result.Error != nil {
		return nil, result.Error
	}
	return &flight, nil
}
