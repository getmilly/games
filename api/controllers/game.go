package controllers

import (
	"net/http"

	"github.com/getmilly/games/pkg/services"
	"github.com/gin-gonic/gin"
)

//GameController ...
type GameController struct {
	gameService services.GameService
}

//NewGameController ...
func NewGameController(gameService services.GameService) *GameController {
	return &GameController{
		gameService: gameService,
	}
}

//RegisterRoutes ...
func (controller *GameController) RegisterRoutes(router *gin.RouterGroup) {
	games := router.Group("games")
	{
		games.POST("", controller.post)
		games.GET("", controller.getAll)
		games.GET(":id", controller.getByID)
		games.DELETE(":id", controller.delete)
	}
}

func (controller *GameController) post(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (controller *GameController) getAll(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (controller *GameController) getByID(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (controller *GameController) delete(c *gin.Context) {
	c.Status(http.StatusOK)
}
