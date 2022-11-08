package controller

import (
	"github.com/SIT-board/board_backend/global"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Engine *gin.Engine
	Config *global.Config
}

func (r *Controller) Init() {
	imageController := &ImageController{
		Group:       r.Engine.Group("/image"),
		ImageConfig: &r.Config.Image,
	}
	imageController.Init()
}
