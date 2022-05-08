package entity

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Balancer struct {
	gorm.Model

	UsedMachines      pq.Int32Array `gorm:"type:integer[]"`
	ConnectedMachines pq.Int32Array `gorm:"type:integer[]"`
}
