// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/jinzhu/gorm"
	"lab3/pkg/controller"
	"lab3/pkg/db/model"
	"lab3/pkg/service"
)

// Injectors from wire.go:

func InitBalancerController(db *gorm.DB) controller.BalancerController {
	balancerModel := model.NewBalancerModel(db)
	balancerService := service.NewBalancerS(balancerModel)
	balancerController := controller.NewBalancerC(balancerService)
	return balancerController
}

func InitMachineController(db *gorm.DB) controller.MachineController {
	machineModel := model.NewMachineModel(db)
	machineService := service.NewMachineS(machineModel)
	machineController := controller.NewMachineC(machineService)
	return machineController
}
