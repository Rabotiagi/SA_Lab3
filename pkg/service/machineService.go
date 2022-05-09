package service

import (
	"lab3/pkg/db/model"
	"lab3/pkg/tools"
	_ "lab3/pkg/tools"
)

type MachineService interface {
	ChangeState(int, int)
	Exists(int) bool
}

type machineService struct {
	model model.MachineModel
}

func NewMachineS(machinesModel model.MachineModel) MachineService {
	return &machineService{
		model: machinesModel,
	}
}

func (ms *machineService) ChangeState(id int, state int) {
	tools.LogError(ms.model.ChangeState(id, state))
}

func (ms *machineService) Exists(id int) bool {
	return ms.model.Exists(id)
}
