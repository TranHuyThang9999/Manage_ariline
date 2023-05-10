package port

import (
	"btl/infrastructure/model"
	"context"
)

// admin
type RepositoryAdmin interface {
	CreateAccountAdmin(ctx context.Context, user *model.Admin) (bool, error)
	LoginAdmin(ctx context.Context, user *model.UserLogin) (bool, error)
	FindByPhoneNumberAdmin(ctx context.Context, phoneNumber string) (*model.Admin, error) // tim kiem
	UpdateAccount(ctx context.Context, phone_numer string, request *model.AdminUpdate) (bool, error)
	DeleteAccount(ctx context.Context, phone_number string) (bool, error)
	FindAccountByForm(ctx context.Context, user *model.UserByForm) ([]*model.User, error)
}
