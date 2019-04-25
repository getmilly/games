package services

import "github.com/getmilly/games/pkg/repositories"

//GameService ...
type GameService interface{}

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
