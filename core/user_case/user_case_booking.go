package user_case

import (
	"btl/infrastructure/model"
	"context"
)

func (tk *RepositoryUserCase) RegisterTicket(ctx context.Context, ticket *model.BookingRequest) (bool, error) {
	status, err := tk.booking.CreateTicket(ctx, ticket)
	if err != nil {
		return false, err
	}
	return status, nil
}
func (tk *RepositoryUserCase) CanCelTicket(ctx context.Context, phone_number string, booking_id string) (bool, error) {
	status, err := tk.booking.CanCelTicket(ctx, phone_number, booking_id)
	if err != nil {
		return false, err
	}
	return status, nil
}
func (tk *RepositoryUserCase) GetAllTicket(ctx context.Context) ([]*model.Booking, error) {
	//redisClient := cache.NewRedisClient("config/config.yaml")
	//cache := cache.NewRedisCache(tk.booking, redisClient, time.Minute*5)
	//tickets, err := cache.GetAllTicket(ctx)
	info_tickets, err := tk.booking.GetAllTicket(ctx)
	if err != nil {
		return nil, err
	}
	return info_tickets, nil
}
func (tk *RepositoryUserCase) GetTicketByPhoneNumber(ctx context.Context, phone_number string) (*model.Booking, error) {
	tichet, err := tk.booking.GetTicketByPhoneNumber(ctx, phone_number)
	if err != nil {
		return nil, err
	}
	return tichet, nil
}
func (tk *RepositoryUserCase) GetStatusTicket(ctx context.Context, phone_number string, booking_id string) (*model.Booking, error) {
	ticket, err := tk.booking.GetStatusTicket(ctx, phone_number, booking_id)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}
func (tk *RepositoryUserCase) GetAllTicketByForm(ctx context.Context, booking model.BookingByForm) ([]*model.Booking, error) {
	ticket, err := tk.booking.GetInforTicketByForm(ctx, booking)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}
