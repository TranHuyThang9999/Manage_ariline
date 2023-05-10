package port

import (
	"btl/infrastructure/model"

	"golang.org/x/net/context"
)

type RepositoryBooking interface {
	//create -- dawng ky ve
	//thong tin ve
	// huy chuyem
	// thanh toan
	CreateTicket(ctx context.Context, ticket *model.BookingRequest) (bool, error)
	CanCelTicket(ctx context.Context, phone_number string, booking_id string) (bool, error)
	GetAllTicket(ctx context.Context) ([]*model.Booking, error)
	GetTicketByPhoneNumber(ctx context.Context, phone_number string) (*model.Booking, error)
	GetStatusTicket(ctx context.Context, phone_number string, booking_id string) (*model.Booking, error)
}
