package services

import (
	"github.com/getmilly/games/pkg/entities"
	"github.com/getmilly/games/pkg/models"
	"github.com/getmilly/games/pkg/repositories"
)

//GameService ...
type GameService interface {
	Create(game models.Game) error
	Find()
	FindByID(gameID string)
	Delete(gameID string)
}

//gameService ...
type gameService struct {
	gameRepository repositories.GameRepository
}

//NewGameService ...
func NewGameService(
	gameRepository repositories.GameRepository,
) GameService {
	return &gameService{
		gameRepository: gameRepository,
	}
}

func (service gameService) Create(game models.Game) error {

	entity := new(entities.Game)

	service.gameRepository.Insert()
}

func (service gameService) Find() {}

func (service gameService) FindByID(gameID string) {}

func (service gameService) Delete(gameID string) {}
