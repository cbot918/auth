package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":5555") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
