package main

import (
	"net/http"
	"x_go_swagger/swagger/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, api.User{
			Name: "今日方知我是我",
			Age:  100,
			Like: "play games",
		})
	})
	r.Run(":9501")
}
