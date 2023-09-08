package main

import (
	v1 "github.com/Seunghoon-Oh/cloud-ml-studio-manager/controller/v1"
	"github.com/Seunghoon-Oh/cloud-ml-studio-manager/data"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.GET("/studios", v1.GetStudios)
	r.POST("/studio", v1.CreateStudio)

	return r
}

func main() {
	data.SetupRedisClient()
	r := setupRouter()
	r.Run(":8082")
}
