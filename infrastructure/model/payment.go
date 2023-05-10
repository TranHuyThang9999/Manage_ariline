package model

import "time"

type Payment struct {
	PaymentID   string    `json:"payment_id"`
	BookingID   string    `json:"booking_id"`
	PaymentTime time.Time `json:"payment_time"`
	Amount      float64   `json:"amount"`
}
