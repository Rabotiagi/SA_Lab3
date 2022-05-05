package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	logError(err)

	server := gin.Default()
	port := os.Getenv("PORT")

	err = server.Run(":" + port)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
