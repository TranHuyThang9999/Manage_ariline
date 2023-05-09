package model

type Flight struct {
	FlightID        string  `json:"flight_id"`
	NameAirline     string  `json:"name_airline"`
	Destination     string  `json:"destination"`
	Departure       string  `json:"departure"`
	DestinationTime string  `json:"destination_time"`
	DepartureTime   string  `json:"departure_time"`
	RemainingSeats  int     `json:"remaining_seats"`
	TicketType      string  `json:"ticket_type"`
	Fare            float64 `json:"fare"`
	Status          string  `json:"status"`
	Name_flight     string  `json:"name_flight"`
}

type FlightByForm struct {
	FlightID        string  `form:"flight_id"`
	NameAirline     string  `form:"name_airline"`
	Destination     string  `form:"destination"`
	Departure       string  `form:"departure"`
	DestinationTime string  `form:"destination_time"`
	DepartureTime   string  `form:"departure_time"`
	RemainingSeats  int     `form:"remaining_seats"` //Số lượng ghế còn lại
	TicketType      string  `form:"ticket_type"`
	Fare            float64 `form:"fare"`
	Status          string  `form:"status"`
	Name_flight     string  `form:"name_flight"`
}
