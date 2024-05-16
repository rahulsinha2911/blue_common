package users_entity

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Sex       string         `json:"sex"`
	Age       int            `json:"age"`
	CreatedAt *time.Time     `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
