package middleware

import (
	"github.com/gin-gonic/gin"
	"lab3/pkg/service"
	"lab3/pkg/tools"
	"strconv"
)

func ValidateId(service service.MachineService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil || !service.Exists(id) {
			tools.AbortRequest(ctx)
		}
	}
}
