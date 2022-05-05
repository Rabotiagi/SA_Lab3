package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"reflect"
)

type Connection struct {
	Dialect, Host  string
	DBPort, DBName string
	User, Password string
	DB             *gorm.DB
}

func (c Connection) ConnectionURL() string {
	connection := reflect.ValueOf(c)
	for i := 0; i < connection.NumField(); i++ {
		if connection.Field(i).Interface() == "" {
			log.Fatal("Empty fields in DB connection config!\n")
		}
	}

	return fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		c.Host,
		c.User,
		c.DBName,
		c.Password,
		c.DBPort,
	)
}

func (c *Connection) Open() *gorm.DB {
	db, err := gorm.Open(c.Dialect, c.ConnectionURL())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (c *Connection) Close() {
	err := c.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
