package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Seunghoon-Oh/cloud-ml-studio-manager/service"
)

func GetStudios(c *gin.Context) {
	data := service.GetStudios()
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func CreateStudio(c *gin.Context) {
	data := service.CreateStudio()
	c.String(http.StatusOK, data)
}
