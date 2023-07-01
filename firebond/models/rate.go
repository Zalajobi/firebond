package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Rates struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string
	CreatedAt time.Time
	Crypto    string
	Fiat      string
	Value     float64
}
