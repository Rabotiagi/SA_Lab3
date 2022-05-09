package middleware

import (
	"github.com/gin-gonic/gin"
	"lab3/pkg/service"
	"net/http"
	"strconv"
)

func ValidateId(service service.MachineService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil || !service.Exists(id) {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
	}
}

func ValidateRequestBody() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
