package domain

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID             string         `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	FirstName      string         `json:"first_name" gorm:"not null"`
	LastName       string         `json:"last_name" gorm:"not null"`
	Login          string         `json:"login" gorm:"not null"`
	PasswordHash   string         `json:"password_hash" gorm:"not null"`
	FullName       string         `json:"full_name"`
	Phone          string         `json:"phone"`
	Pnfl           string         `json:"pnfl"`
	BirthDate      *time.Time     `json:"birth_date"`
	PassportSeries string         `json:"passport_series"`
	PassportNumber string         `json:"passport_number"`
	LegacyCreated  time.Time      `json:"-" gorm:"column:createdAt;autoCreateTime"`
	LegacyUpdated  time.Time      `json:"-" gorm:"column:updatedAt;autoUpdateTime"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
