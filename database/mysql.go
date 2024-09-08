package database

import (
	"log"

	"github.com/vorkey/go-fiber-todo/config"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open(config.GetDbConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Connected to database!")
}

func GetMysqlInstance() *gorm.DB {
	return DB
}
