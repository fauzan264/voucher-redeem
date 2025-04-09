package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID				uuid.UUID	`gorm:"type:char(36);primaryKey"`
	Name			string		`gorm:"type:varchar(100);not null"`
	Email			string		`gorm:"type:varchar(100);not null;uniqueIndex;idx_email"`
	PasswordHash	string		`gorm:"type:varchar(255);not null"`
	TotalPoints		int			`gorm:"type:int;not null;default:0"`
	IsAdmin			bool		`gorm:"type:int;not null;default:0"`
	CreatedAt		time.Time	`gorm:"type:datetime"`
	UpdatedAt		time.Time	`gorm:"type:datetime"`
}