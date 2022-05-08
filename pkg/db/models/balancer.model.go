package model

import (
	"github.com/jinzhu/gorm"
	"lab3/pkg/db/entity"
)

type BalancerModel interface {
	FindAll() []entity.Balancer
}

type balancerModel struct {
	db *gorm.DB
}

func NewBalancerModel(database *gorm.DB) BalancerModel {
	return &balancerModel{db: database}
}

func (bm *balancerModel) FindAll() []entity.Balancer {
	var res []entity.Balancer
	bm.db.Preload("ConnectedMachines").Find(&res)
	return res
}
