package router

import (
	"net/http"
	//generated import

	"Egy2015/cms_go/controller/auth"
	database "Egy2015/cms_go/database/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(db database.DB) (err error) {
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	router.GET("", func(c *gin.Context) {
		db, err := database.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to connect to database",
			})
			return
		}
		defer db.Close()

		c.JSON(http.StatusOK, gin.H{
			"message": "ping-pong!",
		})
	})

	authCtrl := auth.NewAuthController(db.GormDb)
	authRoutes := router.Group("auth")
	{
		authRoutes.POST("/register", authCtrl.Register)
		authRoutes.GET("/login", authCtrl.Login)
	}

	//generated routes
	return router.Run()
}
