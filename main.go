package main

import (
	"net/http"

	v1 "github.com/Seunghoon-Oh/cloud-ml-studio-manager/controller/v1"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "홈페이지 입니다.",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "notebook")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// r.GET("/notebook", func(c *gin.Context) {
	// 	req, err := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	//필요시 헤더 추가 가능
	// 	req.Header.Add("User-Agent", "Crawler")

	// 	// Client객체에서 Request 실행
	// 	client := &http.Client{}
	// 	resp, err := client.Do(req)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer resp.Body.Close()

	// 	// 결과 출력
	// 	bytes, _ := io.ReadAll(resp.Body)
	// 	str := string(bytes) //바이트를 문자열로
	// 	fmt.Println(str)
	// 	c.String(200, "ok")
	// })

	r.GET("/notebook", v1.GetNotebooks)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8082")
}
