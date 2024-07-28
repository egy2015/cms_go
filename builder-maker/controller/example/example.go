package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExampleController struct {
	DB *gorm.DB
}

func NewExampleController(db *gorm.DB) *ExampleController {
	return &ExampleController{DB: db}
}

func (ctrl *ExampleController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "registration successful",
	})
}

func (ctrl *ExampleController) Login(c *gin.Context) {
	// Implement login logic here, interacting with the database using ctrl.DB
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
	})
}
