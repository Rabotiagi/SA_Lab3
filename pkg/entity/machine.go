package entity

import "github.com/jinzhu/gorm"

type Machine struct {
	gorm.Model

	State int
}
