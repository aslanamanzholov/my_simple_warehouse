package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"my_simple_warehouse/pkg/common/models"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Product{})

	return db
}
