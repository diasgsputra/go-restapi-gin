package activitiescontroller

import (
	"encoding/json"
	"net/http"

	"github.com/diasgsputra/go-restapi-gin/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var activities []models.Activities

	models.DB.Find(&activities)
	c.JSON(http.StatusOK, gin.H{"activities": activities})

}

func Show(c *gin.Context) {
	var activities models.Activities
	id := c.Param("id")

	if err := models.DB.First(&activities, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"activities": activities})
}

func Create(c *gin.Context) {

	var activities models.Activities

	if err := c.ShouldBindJSON(&activities); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&activities)
	c.JSON(http.StatusOK, gin.H{"activities": activities})
}

func Update(c *gin.Context) {
	var activities models.Activities
	id := c.Param("id")

	if err := c.ShouldBindJSON(&activities); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&activities).Where("id = ?", id).Updates(&activities).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate activities"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var activities models.Activities

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&activities, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus activities"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
