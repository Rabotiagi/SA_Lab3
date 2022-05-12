//go:build exclude
// +build exclude

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"lab3/pkg/controller"
	"lab3/pkg/db/model"
	"lab3/pkg/service"
)

func InitBalancerController(db *gorm.DB) controller.BalancerController {
	wire.Build(controller.NewBalancerC, service.NewBalancerS, model.NewBalancerModel)
	return nil
}

func InitMachineController(db *gorm.DB) controller.MachineController {
	wire.Build(controller.NewMachineC, service.NewMachineS, model.NewMachineModel)
	return nil
}
