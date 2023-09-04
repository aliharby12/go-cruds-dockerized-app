package main

import (
	"crud-app/inits"
	"crud-app/models"
)

func init() {
	inits.LoadEnvVariables()
	inits.ConnectDB()
}

func main() {
	inits.DB.AutoMigrate(&models.Post{})
	inits.DB.AutoMigrate(&models.User{})
}
