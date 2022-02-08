package utils
/*
* Copyright © 2022 Allan Nava <>
* Created 08/02/2022
* Updated 08/02/2022
*
 */
 import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	if os.Getenv("APP_ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
