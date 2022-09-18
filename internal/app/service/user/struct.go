package user

import "time"

type User struct {
	ID        int64     `json:"id" deepcopier:"ID"`
	Email     string    `json:"email" deepcopier:"email"`
	FullName  string    `json:"full_name" deepcopier:"FullName"`
	Phone     string    `json:"phone" deepcopier:"Phone"`
	Password  string    `json:"password,omitempty" deepcopier:"Password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRes struct {
	AccessToken string    `json:"access_token"`
	ExpiredAt   time.Time `json:"expired_at"`
}

type Res struct {
	ID        int64     `json:"id" deepcopier:"ID"`
	Email     string    `json:"email" deepcopier:"email"`
	FullName  string    `json:"full_name" deepcopier:"FullName"`
	Phone     string    `json:"phone" deepcopier:"Phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
