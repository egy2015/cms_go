package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "registration successful",
	})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	// Implement login logic here, interacting with the database using ctrl.DB
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
	})
}
