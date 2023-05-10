package model

type Admin struct {
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	DateBirth   string `json:"date_birth"`
	NumberCMND  string `json:"number_cmnd"`
	Nationality string `json:"nationality"`
	Language    string `json:"language"`
}
type AdminUpdate struct {
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	DateBirth   string `json:"date_birth"`
	NumberCMND  string `json:"number_cmnd"`
	Nationality string `json:"nationality"`
	Language    string `json:"language"`
}
type AdminLogin struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
