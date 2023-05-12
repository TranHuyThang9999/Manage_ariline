package postGresql

import (
	"btl/infrastructure/model"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type collection struct {
	db *gorm.DB
}

func NewPostGresql(db *gorm.DB) *collection {
	return &collection{
		db: db,
	}
}
func (p *collection) CreateUser(ctx context.Context, user *model.User) (bool, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	user.Password = string(hashedPassword)
	user.UserID = uuid.New().String()
	result := p.db.Create(user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (p *collection) Login(ctx context.Context, user *model.UserLogin) (bool, error) {
	var u model.User
	result := p.db.Where("phone_number = ?", user.PhoneNumber).First(&u)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf("record not found for phone number %s", user.PhoneNumber)
		}
		return false, result.Error
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)); err != nil {
		return false, errors.New("invalid password")
	}
	return true, nil
}

func (p *collection) FindInfo(ctx context.Context, phoneNumber string, password string) (*model.User, error) {
	var u model.User
	result := p.db.Where("phone_number = ?", phoneNumber).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}

	if !u.CheckPassword(password) {
		return nil, errors.New("invalid password")
	}

	return &u, nil
}
func (p *collection) UpdateUser(ctx context.Context, user *model.UserUpdate, name string, phoneNumber string) (bool, error) {
	result := p.db.Model(&model.User{}).Where("user_name = ? AND phone_number = ?", name, phoneNumber).Updates(model.User{
		UserName:    user.UserName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Address:     user.Address,
		DateBirth:   user.DateBirth,
		NumberCMND:  user.NumberCMND,
		Nationality: user.Nationality,
		Language:    user.Language,
		Amount:      user.Amount,
	})
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (p *collection) UpdatePassword(ctx context.Context, phoneNumber string, oldPassword string, newPassword string) (bool, error) {
	var u model.User
	result := p.db.Where("phone_number = ?", phoneNumber).First(&u)
	if result.Error != nil {
		return false, result.Error
	}
	if !u.CheckPassword(oldPassword) {
		return false, errors.New("invalid password")
	}

	err := u.SetPassword(newPassword)
	if err != nil {
		return false, err
	}
	result = p.db.Save(&u)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil

}

//func (p *collection) FindAll(ctx context.Context) ([]*model.User, error) {
//	var users []*model.User
//	result := p.db.Find(&users)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return users, nil
//}

func (p *collection) FindByNumber(ctx context.Context, phoneNumber string) (*model.User, error) {
	var user model.User
	result := p.db.Where("phone_number = ?", phoneNumber).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (p *collection) GetUserById(ctx context.Context, user_id string) (*model.User, error) {
	var user model.User
	result := p.db.Where("user_id = ?", user_id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
