package controller

import (
	"github.com/gin-gonic/gin"
	"lab3/pkg/db/entity"
	"lab3/pkg/service"
	"net/http"
)

type BalancerController interface {
	getAll() []entity.BalancerData
	Config(*gin.Engine)
}

type balancerController struct {
	balancerService service.BalancerService
}

func NewBalancerC(service service.BalancerService) BalancerController {
	return &balancerController{
		balancerService: service,
	}
}

func (bc *balancerController) Config(server *gin.Engine) {
	server.GET("/balancers", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, bc.getAll())
	})
}

func (bc *balancerController) getAll() []entity.BalancerData {
	return bc.balancerService.FindAll()
}
