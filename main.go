package main

import (
	"fmt"
	"log"

	"github.com/alfiqo/disbursement/config"
	"github.com/alfiqo/disbursement/internal/entity"
	"github.com/alfiqo/disbursement/routes"
	"gorm.io/gorm"
)

func main() {
	db := config.DBConnection()
	router := routes.NewResthandler(db)
	DBMigrate(db)
	router.Listen(":8080")
}

// for test only
func DBMigrate(db *gorm.DB) {
	err := db.AutoMigrate(entity.User{}, entity.Wallet{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("succes migrate")
}
