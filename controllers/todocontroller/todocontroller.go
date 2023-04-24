package todocontroller

import (
	"encoding/json"
	"net/http"
	"time"
	"go-restapi-gin/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {

	var todo []models.Todo

	models.DB.Find(&todo)
	c.JSON(http.StatusOK, gin.H{"todo": todo})

}

func GetTodoById(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := models.DB.First(&todo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func CreateTodo(c *gin.Context) {

	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&todo)
	todo.CreatedAt = time.Now()
	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&todo).Where("id = ?", id).Updates(&todo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func DeleteTodo(c *gin.Context) {

	var todo models.Todo

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&todo, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
