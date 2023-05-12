package entities

import "time"

type User struct {
	UserID      string    `json:"user_id"`
	UserName    string    `json:"user_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Address     string    `json:"address"`
	DateOfBirth time.Time `json:"date_birth"`
	NumberCMND  string    `json:"number_cmnd"`
	Nationality string    `json:"nationality"`
	Language    string    `json:"language"`
	Amount      float64   `json:"amount"`
}

type UserLogin struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type Flight struct {
	FlightID        string    `json:"flight_id"`
	NameAirline     string    `json:"name_airline"`
	Destination     string    `json:"destination"`
	Departure       string    `json:"departure"`
	DestinationTime time.Time `json:"destination_time"`
	DepartureTime   time.Time `json:"departure_time"`
	RemainingSeats  int       `json:"remaining_seats"`
	TicketType      string    `json:"ticket_type"`
	Fare            float64   `json:"fare"`
	Status          string    `json:"status"`
}

type Booking struct {
	BookingID     string  `json:"booking_id"`
	UserID        int64   `json:"user_id"`
	FlightID      int64   `json:"flight_id"`
	NumberOfSeats int     `json:"number_of_seats"`
	Amount        float64 `json:"amount"`
}

type Payment struct {
	PaymentID   string    `json:"payment_id"`
	BookingID   string    `json:"booking_id"`
	PaymentTime time.Time `json:"payment_time"`
	Amount      float64   `json:"amount"`
}

type FlightByForm struct {
	FlightID        string    `form:"flight_id"`
	NameAirline     string    `form:"name_airline"`
	Destination     string    `form:"destination"`
	Departure       string    `form:"departure"`
	DestinationTime time.Time `form:"destination_time"`
	DepartureTime   time.Time `form:"departure_time"`
	RemainingSeats  int       `form:"remaining_seats"`
	TicketType      string    `form:"ticket_type"`
	Fare            float64   `form:"fare"`
	Status          string    `form:"status"`
}
