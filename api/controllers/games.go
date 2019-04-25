package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GamesController ...
type GamesController struct {
}

//NewGamesController ...
func NewGamesController() *GamesController {
	return &GamesController{}
}

//RegisterRoutes ...
func (controller *GamesController) RegisterRoutes(router *gin.RouterGroup) {
	games := router.Group("games")
	{
		games.POST("", controller.post)
		games.GET("", controller.getAll)
		games.GET(":id", controller.getByID)
		games.DELETE(":id", controller.delete)
	}
}

func (controller *GamesController) post(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (controller *GamesController) getAll(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (controller *GamesController) getByID(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (controller *GamesController) delete(c *gin.Context) {
	c.Status(http.StatusOK)
}
