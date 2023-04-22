package main

import (
	"github.com/diasgsputra/go-restapi-gin/controllers/activitycontroller"
	"github.com/diasgsputra/go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/activities", activitycontroller.Index)
	r.GET("/api/activity/:id", activitycontroller.Show)
	r.POST("/api/activity", activitycontroller.Create)
	r.PUT("/api/activity/:id", activitycontroller.Update)
	r.DELETE("/api/activity", activitycontroller.Delete)

	r.Run()
}
