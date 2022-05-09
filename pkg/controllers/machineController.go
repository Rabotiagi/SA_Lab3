package controller

import (
	"github.com/gin-gonic/gin"
	"lab3/pkg/db/entity"
	"lab3/pkg/service"
	"lab3/pkg/tools"
	"net/http"
	"strconv"
)

type MachineController interface {
	changeState(int, entity.MachineData) map[string]string
	Config(*gin.Engine)
}
type machineController struct {
	machineService service.MachineService
}

func NewMachineC(service service.MachineService) MachineController {
	return &machineController{
		machineService: service,
	}
}

func (mc *machineController) Config(server *gin.Engine) {
	server.PUT("/machines/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		tools.LogError(err)

		var reqBody entity.MachineData
		err = ctx.BindJSON(&reqBody)
		if err != nil {
			tools.LogError(err)
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"msg":    "invalid input data",
				"status": "rejected",
			})
			return
		}

		ctx.JSON(http.StatusCreated, mc.changeState(id, reqBody))
	})
}

func (mc *machineController) changeState(id int, machine entity.MachineData) map[string]string {
	mc.machineService.ChangeState(id, machine.State)
	return map[string]string{
		"status": "success",
		"msg":    "machine status was changed",
	}
}
