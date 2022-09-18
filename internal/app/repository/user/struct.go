package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64          `json:"id" deepcopier:"ID"`
	Email     string         `json:"email" deepcopier:"email"`
	FullName  string         `json:"full_name" deepcopier:"FullName"`
	Phone     string         `json:"phone" deepcopier:"Phone"`
	Password  string         `json:"password,omitempty" deepcopier:"Password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// TableName ...
func (User) TableName() string {
	return "users"
}
