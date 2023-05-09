package port

import (
	"btl/infra/model"
	"context"
)

// user
type RepositoryUser interface {
	CreateUser(ctx context.Context, user *model.User) (bool, error)
	Login(ctx context.Context, user *model.UserLogin) (bool, error)
	UpdateUser(ctx context.Context, user *model.UserUpdate, user_name string, phoneNumber string) (bool, error)
	UpdatePassword(ctx context.Context, phoneNumber string, oldPassword string, newPassword string) (bool, error)
	FindByNumber(ctx context.Context, phoneNumber string) (*model.User, error) //
}
