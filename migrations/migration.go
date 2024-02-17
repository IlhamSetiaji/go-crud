package main

import (
	"github.com/IlhamSetiaji/go-crud/initializers"
	"github.com/IlhamSetiaji/go-crud/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
