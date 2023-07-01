package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Rates struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	Crypto    string    `gorm:"not null"`
	Fiat      string    `gorm:"not null"`
	Value     float64   `gorm:"not null"`
}
