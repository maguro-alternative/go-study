package db

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Table    string
}

func getDBConfig() DBConfig {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		port = 3306
	}
	return DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Table:    os.Getenv("DB"),
	}
}

func ConnectionDB() (*gorm.DB, error) {
	config := getDBConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", config.User, config.Password, config.Host, config.Port, config.Table)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
