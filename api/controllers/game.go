package controllers

import (
	"fmt"
	"net/http"

	"github.com/getmilly/games/pkg/models"
	"github.com/getmilly/games/pkg/services"
	"github.com/getmilly/grok/api"
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
	var request models.GameRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		api.BindingError(c, err)
		return
	}

	response, err := controller.gameService.Create(request)

	if err != nil {
		api.ResolveError(c, err)
		return
	}

	c.Header("Location", fmt.Sprintf("%s/%s", c.Request.URL.Path, response.ID))
	c.JSON(http.StatusCreated, response)
}

func (controller *GameController) getAll(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (controller *GameController) getByID(c *gin.Context) {
	ID := c.Param("id")

	response, err := controller.gameService.FindByID(ID)

	if err != nil {
		api.ResolveError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (controller *GameController) delete(c *gin.Context) {
	c.Status(http.StatusOK)
}
