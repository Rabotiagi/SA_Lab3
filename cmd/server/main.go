package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"lab3/pkg/db"
	"log"
	"os"
)

var err = godotenv.Load(".env")

var (
	port     = os.Getenv("PORT")
	dialect  = os.Getenv("DIALECT")
	host     = os.Getenv("HOST")
	dbPort   = os.Getenv("DBPORT")
	dbName   = os.Getenv("DBNAME")
	user     = os.Getenv("DBUSER")
	password = os.Getenv("PASSWORD")
)

func main() {
	logError(err)
	dbConnection := db.Connection{
		Dialect:  dialect,
		Host:     host,
		DBPort:   dbPort,
		DBName:   dbName,
		User:     user,
		Password: password,
	}

	err = dbConnection.Open()
	defer dbConnection.Close()
	logError(err)
	err = dbConnection.AutoFill()
	logError(err)

	server := gin.Default()
	err = server.Run(":" + port)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
