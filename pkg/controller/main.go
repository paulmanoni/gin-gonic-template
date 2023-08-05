package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainController interface {
	Home(c *gin.Context)
}

type mainController struct{}

func (m mainController) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func NewMainController() MainController {
	return &mainController{}
}
