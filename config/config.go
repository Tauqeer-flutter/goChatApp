package config

import (
	"goChatApp/domain"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB     *gorm.DB
	SERVER string
	PORT   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := connectDb()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		DB:     db,
		SERVER: os.Getenv("SERVER"),
		PORT:   os.Getenv("PORT"),
	}
}

func connectDb() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.User{}, &domain.Group{}, &domain.Member{}, &domain.Chat{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
