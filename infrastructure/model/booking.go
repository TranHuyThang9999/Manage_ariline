package model

type Booking struct {
	BookingID       string  `json:"booking_id"`
	UserID          string  `json:"user_id"`
	FlightID        string  `json:"flight_id"`
	NumberOfSeats   int     `json:"number_of_seats"`
	Amount          float64 `json:"amount"`
	UserName        string  `json:"user_name"`
	PhoneNumber     string  `json:"phone_number"`
	Address         string  `json:"address"`
	NumberCMND      string  `json:"number_cmnd"`
	Nationality     string  `json:"nationality"`
	Language        string  `json:"language"`
	NameAirline     string  `json:"name_airline"`
	Destination     string  `json:"destination"`
	Departure       string  `json:"departure"`
	DestinationTime string  `json:"destination_time"`
	DepartureTime   string  `json:"departure_time"`
	TicketType      string  `json:"ticket_type"`
	Fare            float64 `json:"fare"`
	Status          string  `json:"status"`
	NameFlight      string  `json:"name_flight"`
	RegisterTime    string  `json:"register_time"`
}
type BookingRequest struct {
	BookingID       string `json:"booking_id"`
	FlightID        string `json:"flight_id"`
	PhoneNumber     string `json:"phone_number"` //header
	NumberOfSeats   int    `json:"number_of_seats"`
	TicketType      string `json:"ticket_type"`
	NameAirline     string `json:"name_airline"`
	Destination     string `json:"destination"`
	Departure       string `json:"departure"`
	DestinationTime string `json:"destination_time"`
	DepartureTime   string `json:"departure_time"`
}
type BookingByForm struct {
	BookingID       string  `form:"booking_id"`
	UserID          string  `form:"user_id"`
	FlightID        string  `form:"flight_id"`
	NumberOfSeats   int     `form:"number_of_seats"`
	Amount          float64 `form:"amount"`
	UserName        string  `form:"user_name"`
	PhoneNumber     string  `form:"phone_number"`
	Address         string  `form:"address"`
	NumberCMND      string  `form:"number_cmnd"`
	Nationality     string  `form:"nationality"`
	Language        string  `form:"language"`
	NameAirline     string  `form:"name_airline"`
	Destination     string  `form:"destination"`
	Departure       string  `form:"departure"`
	DestinationTime string  `form:"destination_time"`
	DepartureTime   string  `form:"departure_time"`
	TicketType      string  `form:"ticket_type"`
	Fare            float64 `form:"fare"`
	Status          string  `form:"status"`
	NameFlight      string  `form:"name_flight"`
	RegisterTime    string  `:"register_time"`
}
