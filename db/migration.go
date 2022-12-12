package db

import (
	"fmt"
	 "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/lib/pq"
)

const (
	DBURL = "postgres://root:123@localhost:5432/Users?sslmode=disable"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DBURL),&gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect")
	}
	db.AutoMigrate(&User{})
	return db
}
