package main

import (
	"github.com/pondelion/dbeservice/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	engine.Use(cors.New(config))

	networkEngine := engine.Group("/network")
	{
		networkEngine.POST("/start_tcpdump", controller.StartTcpdump)
		networkEngine.POST("/stop_tcpdump", controller.StopTcpdump)
	}
	engine.Run()
}
