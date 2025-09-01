package config

import (
	"fmt"
	"log"
	"wallet-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB *gorm.DB

	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func NewDBConfig() DBConfig {
	return DBConfig{
		DB:       nil,
		Host:     "",
		User:     "",
		Password: "",
		Name:     "",
		Port:     "",
	}
}

// get input from config struct
func (c *DBConfig) InitDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		c.Host,
		c.User,
		c.Password,
		c.Name,
		c.Port,
	)

	c.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}

	log.Println("Database connected successfully")
}

func (c *DBConfig) Migrate() {
	err := c.DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database.", err)
	}
	log.Println("Database migration completed")
}
