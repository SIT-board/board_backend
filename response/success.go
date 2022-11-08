package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}
