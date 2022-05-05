package entity

import "github.com/jinzhu/gorm"

type Balancer struct {
	gorm.Model

	Id                int
	UsedMachines      []int
	ConnectedMachines []int
}
