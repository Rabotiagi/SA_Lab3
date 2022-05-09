package controller

import (
	"github.com/gin-gonic/gin"
	"lab3/pkg/middleware"
	"lab3/pkg/service"
	"lab3/pkg/tools"
	"net/http"
	"strconv"
)

type MachineController interface {
	changeState(int, int) map[string]string
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
	server.PUT("/machines/:id", middleware.ValidateId(mc.machineService), func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		tools.LogError(err)
		state, valid := tools.ValidateReqBody(ctx)

		if valid {
			ctx.JSON(http.StatusCreated, mc.changeState(id, state))
		}
	})
}

func (mc *machineController) changeState(id int, state int) map[string]string {
	mc.machineService.ChangeState(id, state)
	return map[string]string{
		"status": "success",
		"msg":    "machine status was changed",
	}
}
