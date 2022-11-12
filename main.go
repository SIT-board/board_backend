package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/SIT-board/board_backend/controller"
	"github.com/SIT-board/board_backend/global"
)

func main() {
	r := gin.Default()

	config := global.GetDefaultConfig()
	(&controller.Controller{
		Engine: r,
		Config: config,
	}).Init()

	server := config.Server
	log.Fatal(r.Run(fmt.Sprintf("%s:%d", server.ListenIp, server.ListenPort)))
}
