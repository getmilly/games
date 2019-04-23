package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AppController ...
type AppController struct {
}

//NewAppController ...
func NewAppController() *AppController {
	return &AppController{}
}

//RegisterRoutes ...
func (controller *AppController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/app", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
