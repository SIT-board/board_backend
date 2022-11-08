package response

import "github.com/gin-gonic/gin"

func ResponseError(ctx *gin.Context, status int, err error) {
	er := ErrorBody{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

type ErrorBody struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
