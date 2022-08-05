package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInitialization() {
	var err error

	//const URI string = "root:@tcp(127.0.0.1:3306)/gofiber?charset=utf8mb4&parseTime=True&loc=Local"
	const POSTGRESTDB string = "postgresql://aksesaja:Aksesaja123!@db-aksesaja.cnqous4emciz.ap-southeast-1.rds.amazonaws.com:5432/gofiber"
	DB, err = gorm.Open(postgres.Open(POSTGRESTDB), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Database connected")
}
