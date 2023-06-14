package user_case

import (
	"btl/infrastructure/model"
	"context"
)

func (p *RepositoryUserCase) CreateAccountAdmin(ctx context.Context, user *model.Admin) (bool, error) {
	status, err := p.admin_user_case.CreateAccountAdmin(ctx, user)
	if err != nil {
		return false, err
	}
	return status, nil
}

func (p *RepositoryUserCase) LoginAdmin(ctx context.Context, user *model.UserLogin) (bool, error) {
	status, err := p.admin_user_case.LoginAdmin(ctx, user)
	if err != nil {
		return false, err
	}
	if !status {
		return false, nil
	}
	return status, nil
}
func (p *RepositoryUserCase) FindByPhoneNumberAdmin(ctx context.Context, phoneNumber string) (*model.Admin, error) {
	admin, err := p.admin_user_case.FindByPhoneNumberAdmin(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}
	return admin, err
}
func (p *RepositoryUserCase) FindByForm(ctx context.Context, user *model.UserByForm) ([]*model.User, error) {
	users, err := p.admin_user_case.FindAccountByForm(ctx, user)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, nil
	}
	return users, nil
}
