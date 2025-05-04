package db

import (
	"fmt"
	"log"
	"todoapp/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=admin password=secret dbname=todo_db port=5432 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error database connect: ", err)
	}

	err = DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal("migrate error: ", err)
	}
	fmt.Println("Connected and migrate.")
}
