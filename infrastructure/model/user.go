package model

import "golang.org/x/crypto/bcrypt"

type UserUpdate struct {
	UserName    string  `json:"user_name"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Address     string  `json:"address"`
	DateBirth   string  `json:"date_birth"`
	NumberCMND  string  `json:"number_cmnd"`
	Nationality string  `json:"nationality"`
	Language    string  `json:"language"`
	Amount      float64 `json:"amount"`
}
type User struct {
	UserID      string  `json:"user_id"`
	UserName    string  `json:"user_name"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Address     string  `json:"address"`
	DateBirth   string  `json:"date_birth"`
	NumberCMND  string  `json:"number_cmnd"`
	Nationality string  `json:"nationality"`
	Language    string  `json:"language"`
	Amount      float64 `json:"amount"`
}

type UserLogin struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = ""
	u.Password = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type UserByForm struct {
	UserID      string `form:"user_id"`
	UserName    string `form:"user_name"`
	PhoneNumber string `form:"phone_number"`
	Email       string `form:"email"`
	Address     string `form:"address"`
	DateBirth   string `form:"date_birth"`
	NumberCMND  string `form:"number_cmnd"`
	Nationality string `form:"nationality"`
	Language    string `form:"language"`
}
