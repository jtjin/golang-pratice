package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jtjin/golang-pratice/controller"
	"github.com/jtjin/golang-pratice/middlewares"
	"github.com/jtjin/golang-pratice/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	// Logging to a file.
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// Logging to a file AND console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// server := gin.Default()
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middlewares.Logger())
	server.Use(middlewares.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	server.Run(":8080")
}
