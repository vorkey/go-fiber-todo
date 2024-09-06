package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type databaseConfig struct {
	user     string
	password string
	host     string
	port     string
	dbName   string
}

func init() {
	log.Fatalln(godotenv.Load())
}

func DatabaseInit() gorm.Dialector {
	dbConfig := databaseConfig{
		user:     os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		dbName:   os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf(
		"user=%s:password=%s@tcp(host=%s:port=%s)/dbname=%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.user,
		dbConfig.password,
		dbConfig.host,
		dbConfig.port,
		dbConfig.dbName,
	)

	return mysql.Open(dsn)
}
