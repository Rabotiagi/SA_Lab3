package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateReqBody(ctx *gin.Context) (int, bool) {
	flag := true
	var reqBody map[string]int
	err := ctx.BindJSON(&reqBody)

	for key, val := range reqBody {
		if key == "state" && (val == 0 || val == 1) {
			flag = false
			break
		}
	}

	if err != nil || flag {
		AbortRequest(ctx)
	}

	return reqBody["state"], err == nil && !flag
}

func AbortRequest(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
		"msg":    "invalid input data",
		"status": http.StatusBadRequest,
	})
}
