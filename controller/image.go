package controller

import (
	"net/http"

	"github.com/SIT-board/board_backend/global"
	"github.com/SIT-board/board_backend/response"
	"github.com/SIT-board/board_backend/service"
	"github.com/gin-gonic/gin"
)

type ImageController struct {
	Group       *gin.RouterGroup
	ImageConfig *global.ImageConfig

	imageService *service.ImageService
}

func (r *ImageController) Init() {
	r.Group.POST("/upload", r.Upload)
	r.imageService = &service.ImageService{
		RawDirectory:       r.ImageConfig.RawDirectory,
		ThumbnailDirectory: r.ImageConfig.ThumbnailDirectory,
	}
}

func (r *ImageController) Upload(c *gin.Context) {
	file, err := c.FormFile("image")

	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, err)
	}
	filename, err := r.imageService.SaveImage(c, file)
	if err != nil {
		response.ResponseError(c, http.StatusBadRequest, err)
	}
	response.ResponseSuccess(c, filename)
}
