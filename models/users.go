package models

import "github.com/google/uuid"

type Users struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	Email     string    `json:"email" gorm:"unique"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}
