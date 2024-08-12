package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"x_go_swagger/api"
)

var localUsers []api.User

func main() {
	r := gin.Default()
	//支持swagger跨域访问
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true, //本地测试，所以运行所有域名，ip跨域访问
	}))

	r.GET("/user", func(c *gin.Context) {
		userName := c.Param("name")
		for _, user := range localUsers {
			if user.Name == userName {
				c.JSON(http.StatusOK, user)
				return
			}
		}

		c.JSON(http.StatusOK, api.ErrResponse{
			Code:    400,
			Message: "查不到该用户",
		})
	})

	log.Fatal(r.Run(":9501"))
}
