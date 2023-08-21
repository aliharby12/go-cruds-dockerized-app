package inits

import (
	"log"

	"github.com/joho/godotenv"
)

// start with capital letter because we will
// use it in another file
func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading godotenv")
	}
}
