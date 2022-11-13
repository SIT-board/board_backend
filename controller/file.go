package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SIT-board/board_backend/global"
	"github.com/SIT-board/board_backend/response"
	"github.com/SIT-board/board_backend/service"
)

type AttachmentController struct {
	Group   *gin.RouterGroup
	Config  *global.FileConfig
	service *service.FileService
}

func (controller *AttachmentController) Init() {
	controller.service = &service.FileService{Directory: controller.Config.Directory}
	controller.Group.POST("/upload", controller.Upload)
	controller.Group.StaticFS("/download", http.Dir(controller.Config.Directory))

}

type uploadResp struct {
	Url string `json:"url"`
}

func (controller *AttachmentController) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	fileName, err := controller.service.SaveFile(ctx, file)
	if err != nil {
		response.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	resp := uploadResp{Url: fmt.Sprintf("%s/file/download/%s", controller.Config.Host, fileName)}

	response.ResponseSuccess(ctx, resp)
}
