package utils

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"github.com/nogsantos/noggo/domain"
)

// https://v1.gorm.io/docs/connecting_to_the_database.html

func ConnectDB() *gorm.DB {
	env_error := godotenv.Load()

	if env_error != nil {
		log.Fatalf("Error to load env")
		panic(env_error)
	}

	dns := os.Getenv("dns")
	dbType := os.Getenv("dbType")

	db, db_error := gorm.Open(dbType, dns)

	if db_error != nil {
		log.Fatalf("Error to setup the database")
		panic(db_error)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	// defer db.Close()

	db.AutoMigrate(&domain.User{})

	return db

}