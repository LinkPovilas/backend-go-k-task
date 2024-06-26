package main

import (
	"log"

	db "github.com/LinkPovilas/backend-go-k-task/database"
	"github.com/LinkPovilas/backend-go-k-task/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()

	r := routes.SetupRouter()
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
