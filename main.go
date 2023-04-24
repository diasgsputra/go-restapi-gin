package main

import (
	"go-restapi-gin/controllers/activitycontroller"
	"go-restapi-gin/controllers/todocontroller"
	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/activity-groups", activitycontroller.Index)
	r.GET("/activity-groups/:id", activitycontroller.Show)
	r.POST("/activity-groups", activitycontroller.Create)
	r.PATCH("/activity-groups/:id", activitycontroller.Update)
	r.DELETE("/activity-groups/:id", activitycontroller.Delete)

	r.GET("/todo-items", todocontroller.GetTodo)
	r.GET("/todo-items/:id", todocontroller.GetTodoById)
	r.POST("/todo-items", todocontroller.CreateTodo)
	r.PATCH("/todo-items/:id", todocontroller.UpdateTodo)
	r.DELETE("/todo-items/:id", todocontroller.DeleteTodo)

	r.Run(":3030")
}
