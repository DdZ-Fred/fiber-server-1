package main

import (
	"log"

	fiberserver1 "github.com/DdZ-Fred/fiber-server-1"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	fiberserver1.Run()
}
