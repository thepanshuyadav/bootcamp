package main

import (
	"ecommerce/config"
	"ecommerce/models"
	"ecommerce/routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	// Connect to database
	config.DB, err = gorm.Open("mysql", config.DBUrl(config.DBSetup()))

	if err != nil {
		fmt.Println(err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Product{})

	r := routes.SetUpRouter()
	//running
	r.Run()

}
