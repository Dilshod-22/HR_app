package main

import (
	"HR/internal/delivery/http"
	"HR/internal/infrastructure"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env fayl topilmadi, tizim env o'zgaruvchilari ishlatiladi")
	}

	infrastructure.ConnectDB()
	r := http.EmployeeRoute()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
