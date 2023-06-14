package postGresql

import (
	"btl/infrastructure/model"
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (p *collection) CreateAccountAdmin(ctx context.Context, user *model.Admin) (bool, error) {

	user.UserID = uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	user.Password = string(hashedPassword)

	result := p.db.Create(user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
func (p *collection) UpdateAccount(ctx context.Context, phone_numer string, request *model.AdminUpdate) (bool, error) {
	result := p.db.Model(&model.Admin{}).Where("phone_number = ?", phone_numer).Updates(request)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
func (p *collection) DeleteAccount(ctx context.Context, phone_number string) (bool, error) {
	result := p.db.Delete(&model.User{}, "phone_number = ?", phone_number)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, errors.New("account not found")
	}
	return true, nil
}
func (p *collection) LoginAdmin(ctx context.Context, user *model.UserLogin) (bool, error) {
	var admin model.Admin
	result := p.db.Where("phone_number = ?", user.PhoneNumber).First(&admin)
	if result.Error != nil {
		return false, result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(user.Password))
	if err != nil {
		return false, err
	}
	return true, nil
}
func (p *collection) FindByPhoneNumberAdmin(ctx context.Context, phoneNumber string) (*model.Admin, error) {
	var user model.Admin
	result := p.db.Where("phone_number = ?", phoneNumber).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (p *collection) FindAccountByForm(ctx context.Context, user *model.UserByForm) ([]*model.User, error) {
	users := make([]*model.User, 0)
	result := p.db.Where(model.User{
		UserID:      user.UserID,
		UserName:    user.UserName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Address:     user.Address,
		DateBirth:   user.DateBirth,
		NumberCMND:  user.NumberCMND,
		Nationality: user.Nationality,
		Language:    user.Language,
	}).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return users, nil
}
