package services

import (
	"github.com/getmilly/games/pkg/mappings"
	"github.com/getmilly/games/pkg/models"
	"github.com/getmilly/games/pkg/repositories"
)

//GameService ...
type GameService interface {
	Create(game models.GameRequest) (*models.GameResponse, error)
	Find()
	FindByID(gameID string) (*models.GameResponse, error)
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

func (service gameService) Create(game models.GameRequest) (*models.GameResponse, error) {
	entity := mappings.GameRequestModelToGameEntity(&game)

	err := service.gameRepository.Insert(entity)

	response := mappings.GameEntityToGameResponseModel(entity)

	return response, err
}

func (service gameService) Find() {}

func (service gameService) FindByID(ID string) (*models.GameResponse, error) {
	entity, err := service.gameRepository.FindByID(ID)

	if err != nil {
		return nil, err
	}

	response := mappings.GameEntityToGameResponseModel(entity)

	return response, nil
}

func (service gameService) Delete(ID string) {}
