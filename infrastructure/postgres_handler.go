package infrastructure

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(conString string) *gorm.DB {
	//"host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database : %s", err.Error()))
	}

	return db
}
