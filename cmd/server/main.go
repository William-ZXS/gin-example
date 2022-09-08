package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	//app.RedirectTrailingSlash = false
	//app.RedirectFixedPath = true
	//app.HandleMethodNotAllowed = true
	app.GET("/ping", func(c *gin.Context) {
		fmt.Println(c.Request.RemoteAddr)
		fmt.Println(c.ClientIP())
		c.JSON(200, "pong")
	})
	app.Run(":9000")
}
