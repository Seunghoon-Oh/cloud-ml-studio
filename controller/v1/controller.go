package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Seunghoon-Oh/cloud-ml-studio-manager/service"
)

func GetStudios(c *gin.Context) {
	data := service.GetStudios()
	println("Response: " + data)
	c.String(http.StatusOK, data)
}
