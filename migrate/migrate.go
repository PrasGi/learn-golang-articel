package main

import (
	"github.com/PrasGi/learn-golang/initializers"
	"github.com/PrasGi/learn-golang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Articel{})
	initializers.DB.AutoMigrate(&models.User{})
}
