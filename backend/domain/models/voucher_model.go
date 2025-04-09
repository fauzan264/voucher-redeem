package models

import (
	"time"

	"github.com/google/uuid"
)

type Voucher struct {
	ID				uuid.UUID	`gorm:"type:char(36);primaryKey"`
	BrandID			uuid.UUID	`gorm:"type:char(36);not null"`
	Code			string		`gorm:"type:varchar(50);uniqueIndex"`
	Name			string		`gorm:"type:varchar(100);not null"`
	Description		string		`gorm:"type:text"`
	CostInPoints	int			`gorm:"type:int;not null"`
	Stock			int			`gorm:"type:int"`
	ValidFrom		time.Time	`gorm:"type:datetime"`
	ValidUntil		time.Time	`gorm:"type:datetime"`
	CreatedAt		time.Time	`gorm:"type:datetime"`
	UpdatedAt		time.Time	`gorm:"type:datetime"`

	Brand			Brand		`gorm:"foreignKey:BrandID;references:ID"`
}