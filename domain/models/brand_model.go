package models

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	ID				uuid.UUID	`gorm:"type:char(36);primaryKey"`
	Name			string		`gorm:"type:varchar(100);not null"`
	Description		string		`gorm:"type:varchar(100);not null"`
	CreatedAt		time.Time	`gorm:"type:datetime"`
	UpdatedAt		time.Time	`gorm:"type:datetime"`
}