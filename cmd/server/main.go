package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"lab3/pkg/db"
	"lab3/pkg/tools"
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
	tools.LogError(err)
	dbConnection := db.Connection{
		Dialect:  dialect,
		Host:     host,
		DBPort:   dbPort,
		DBName:   dbName,
		User:     user,
		Password: password,
	}

	tools.LogError(dbConnection.Open())
	defer dbConnection.Close()
	tools.LogError(dbConnection.AutoFill())

	bc := InitBalancerController(dbConnection.DB)
	mc := InitMachineController(dbConnection.DB)

	server := gin.Default()
	bc.Config(server)
	mc.Config(server)
	tools.LogError(server.Run(":" + port))
}
