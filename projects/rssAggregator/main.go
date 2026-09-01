package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not defined in .env variables")
	}

	fmt.Println("PORT:", portString)
}
