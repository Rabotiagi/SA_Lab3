package model

import (
	"github.com/jinzhu/gorm"
	"lab3/pkg/db/entity"
)

type MachineModel interface {
	ChangeState(int, int) error
	findAll() []entity.Machine
	Exists(int) bool
}

type machineModel struct {
	db *gorm.DB
}

func NewMachineModel(database *gorm.DB) MachineModel {
	return &machineModel{db: database}
}

func (mm *machineModel) findAll() []entity.Machine {
	var res []entity.Machine
	mm.db.Find(&res)
	return res
}

func (mm *machineModel) ChangeState(id int, state int) error {
	res := mm.db.Model(&entity.Machine{Id: id}).Update("state", state)
	return res.Error
}

func (mm *machineModel) Exists(id int) bool {
	var exists entity.Machine
	mm.db.Table("machines").
		Where("id = ?", id).
		Find(&exists)
	return exists.Id != 0
}
