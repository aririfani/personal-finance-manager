package user

import "time"

type User struct {
	ID        int64     `json:"id" deepcopier:"ID"`
	Email     string    `json:"email" deepcopier:"email"`
	FullName  string    `json:"full_name" deepcopier:"UserName"`
	Phone     string    `json:"phone" deepcopier:"Phone"`
	Password  string    `json:"password,omitempty" deepcopier:"Password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRes struct {
	AccessToken string    `json:"access_token"`
	ExpiredAt   time.Time `json:"expired_at"`
}
