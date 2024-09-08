package migration

import (
	"fmt"
	"log"

	"github.com/vorkey/go-fiber-todo/database"
	"github.com/vorkey/go-fiber-todo/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
