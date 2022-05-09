package service

import (
	"lab3/pkg/db/entity"
	"lab3/pkg/db/model"
	"lab3/pkg/tools"
)

type BalancerService interface {
	FindAll() []entity.BalancerData
}

type balancerService struct {
	model model.BalancerModel
}

func NewBalancerS(balancerModel model.BalancerModel) BalancerService {
	return &balancerService{
		model: balancerModel,
	}
}

func (bs *balancerService) FindAll() []entity.BalancerData {
	var res []entity.BalancerData
	balancers, err := bs.model.FindAll()
	tools.LogError(err)

	for _, bal := range balancers {
		balData := entity.BalancerData{
			Id:                 bal.Id,
			TotalMachinesCount: len(bal.ConnectedMachines),
			UsedMachines:       []int{},
		}

		for _, machine := range bal.ConnectedMachines {
			if machine.State == 1 {
				balData.UsedMachines = append(balData.UsedMachines, machine.Id)
			}
		}

		res = append(res, balData)
	}

	return res
}
