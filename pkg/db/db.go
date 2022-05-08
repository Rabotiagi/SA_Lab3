package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"lab3/pkg/db/entity"
	"log"
	"reflect"
)

type Connection struct {
	Dialect, Host  string
	DBPort, DBName string
	User, Password string
	DB             *gorm.DB
}

func (c *Connection) connectionURL() string {
	connection := reflect.ValueOf(*c)
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

func (c *Connection) Open() error {
	db, err := gorm.Open(c.Dialect, c.connectionURL())
	c.DB = db

	return err
}

func (c *Connection) Close() error {
	err := c.DB.Close()
	return err
}

func (c *Connection) AutoFill() error {
	if c.DB == nil {
		return errors.New("db connection is not established")
	}

	c.DB.Where("1 = 1").Unscoped().Delete(&entity.Machine{})
	c.DB.Where("1 = 1").Unscoped().Delete(&entity.Balancer{})

	c.DB.AutoMigrate(&entity.Balancer{})
	c.DB.AutoMigrate(&entity.Machine{})

	machine1 := entity.Machine{State: 0}
	machine2 := entity.Machine{State: 1}
	machine3 := entity.Machine{State: 1}
	machine4 := entity.Machine{State: 0}
	c.DB.Create(&machine1)
	c.DB.Create(&machine2)
	c.DB.Create(&machine3)
	c.DB.Create(&machine4)

	balancer1 := entity.Balancer{
		ConnectedMachines: []entity.Machine{machine1, machine2},
	}
	balancer2 := entity.Balancer{
		ConnectedMachines: []entity.Machine{machine3, machine4},
	}
	c.DB.Create(&balancer1)
	c.DB.Create(&balancer2)

	return nil
}
