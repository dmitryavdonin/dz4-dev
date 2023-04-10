package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username,omitempty"`
	FirstName string    `gorm:"not null" json:"firstName,omitempty"`
	LastName  string    `gorm:"not null" json:"lastName,omitempty"`
	Email     string    `gorm:"not null" json:"email,omitempty"`
	Phone     string    `gorm:"not null" json:"phone,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateUserRequest struct {
	Username  string    `json:"username"  binding:"required"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateUser struct {
	Username  string    `json:"username,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
