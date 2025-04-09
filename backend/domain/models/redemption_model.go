package models

import (
	"time"

	"github.com/google/uuid"
)

type Redemption struct {
	ID				uuid.UUID	`gorm:"type:char(36);primaryKey"`
	CustomerID		uuid.UUID	`gorm:"type:char(36);not null"`
	Code 			string  	`gorm:"type:varchar(100);unique"`
	PointUsed		int			`gorm:"type:int"`
	ReedemedAt		time.Time	`gorm:"type:datetime"`

	Customer		User		`gorm:"foreignKey:CustomerID;references:ID"`
}

type RedemptionItem struct {
	ID				uuid.UUID	`gorm:"type:char(36);primaryKey"`
	RedemptionID	uuid.UUID	`gorm:"type:char(36);not null"`
	VoucherID		uuid.UUID	`gorm:"type:char(36);not null"`
	Quantity		int			`gorm:"type:int;not null"`
	SubTotalPoint	int			`gorm:"type:int;not null"`

	Redemption		Redemption	`gorm:"foreignKey:RedemptionID;references:ID"`
	Voucher			Voucher		`gorm:"foreignKey:VoucherID;references:ID"`
}