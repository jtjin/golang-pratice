package main

import (
	"io"
	"net/http"
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

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			video, err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, video)
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}
