package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/SIT-board/board_backend/controller"
	"github.com/SIT-board/board_backend/global"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "conf", "config.yaml", "config path ,ag : go run main -f config.yaml")
}

func main() {
	flag.Parse()
	r := gin.Default()
	config := global.GetConfig(configFile)
	(&controller.Controller{
		Engine: r,
		Config: config,
	}).Init()

	server := config.Server
	log.Fatal(r.Run(fmt.Sprintf("%s:%d", server.ListenIp, server.ListenPort)))
}
