package model

import (
	"github.com/jinzhu/gorm"
	"lab3/pkg/db/entity"
)

type BalancerModel struct {
}

func (bm *BalancerModel) FindAll(db *gorm.DB) []entity.Balancer {
	var balancers []entity.Balancer
	db.Preload("ConnectedMachines").Find(&balancers)
	return balancers
}
