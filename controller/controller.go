package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/SIT-board/board_backend/global"
)

type Controller struct {
	Engine *gin.Engine
	Config *global.Config
}

func (controller *Controller) Init() {
	attachment := &AttachmentController{
		Group:  controller.Engine.Group("/file"),
		Config: &controller.Config.File,
	}
	attachment.Init()
}
