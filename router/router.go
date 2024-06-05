package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() (err error) {
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping-pong!",
		})
	})
	return router.Run()
}
