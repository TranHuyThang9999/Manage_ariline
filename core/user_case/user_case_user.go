package user_case

import (
	"btl/core/port"
	"btl/infra/model"
	"context"
)

type RepositoryUserCase struct {
	user_case       port.RepositoryUser
	admin_user_case port.RepositoryAdmin
	flight_use_case port.RepositoryFlight
	booking         port.RepositoryBooking
}

func NewUserCaseUser(user port.RepositoryUser, admin port.RepositoryAdmin, flight port.RepositoryFlight, booking port.RepositoryBooking) *RepositoryUserCase {
	return &RepositoryUserCase{
		user_case:       user,
		admin_user_case: admin,
		flight_use_case: flight,
		booking:         booking,
	}
}

func (u *RepositoryUserCase) CreateAccountUser(ctx context.Context, user *model.User) (bool, error) {

	status, err := u.user_case.CreateUser(ctx, user)
	if err != nil {
		return false, err
	}
	return status, nil
}

func (u *RepositoryUserCase) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*model.User, error) {
	user, err := u.user_case.FindByNumber(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *RepositoryUserCase) LoginUser(ctx context.Context, user *model.UserLogin) (bool, error) {
	reponse_user, err := u.user_case.Login(ctx, user)
	if err != nil {
		return false, nil
	}
	if !reponse_user {
		return false, nil
	}
	return reponse_user, nil
}
func (u *RepositoryUserCase) UpdateUser(ctx context.Context, user *model.UserUpdate, user_name string, phoneNumber string) (bool, error) {
	status, err := u.user_case.UpdateUser(ctx, user, user_name, phoneNumber)
	if err != nil {
		return false, nil
	}
	return status, nil
}
func (u *RepositoryUserCase) UpdatePassword(ctx context.Context, phoneNumber string, oldPassword string, newPassword string) (bool, error) {
	status, err := u.user_case.UpdatePassword(ctx, phoneNumber, oldPassword, newPassword)
	if err != nil {
		return false, nil
	}
	return status, nil
}

//func (u *RepositoryUserCase) FindALlUser(ctx context.Context) ([]*model.User, error) {
//	users, err := u.user_case.FindAll(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return users, nil
//}
