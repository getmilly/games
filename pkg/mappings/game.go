package mappings

import (
	"github.com/getmilly/games/pkg/entities"
	"github.com/getmilly/games/pkg/models"
)

//GameModelToGameEntity ...
func GameModelToGameEntity(model *models.Game) *entities.Game {
	return &entities.Game{
		Name: model.Name,
	}
}
