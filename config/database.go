package config

import (
	"fmt"
	"log"
	"os"

	"github.com/fauzan264/voucher-redeem/domain/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	var(
		DBHost 		= os.Getenv("DB_HOST")
		DBPort 		= os.Getenv("DB_PORT")
		DBName 		= os.Getenv("DB_NAME")
		DBUser 		= os.Getenv("DB_USER")
		DBPassword	= os.Getenv("DB_PASSWORD")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort,DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Brand{},
		&models.Redemption{},
		&models.RedemptionItem{},
		&models.Voucher{},	
	)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}

	return db
}