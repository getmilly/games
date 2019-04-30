package mappings

import (
	"time"

	"github.com/getmilly/games/pkg/entities"
	"github.com/getmilly/games/pkg/models"
)

//GameRequestModelToGameEntity ...
func GameRequestModelToGameEntity(model *models.GameRequest) *entities.Game {
	createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	startAt, _ := time.Parse(time.RFC3339, model.StartAt.Format(time.RFC3339))
	endAt, _ := time.Parse(time.RFC3339, model.EndAt.Format(time.RFC3339))

	return &entities.Game{
		Name:        model.Name,
		Description: model.Description,
		StartAt:     startAt,
		EndAt:       endAt,
		URLImage:    model.URLImage,
		Active:      true,
		CreatedAt:   createdAt,
	}
}

//GameEntityToGameResponseModel ...
func GameEntityToGameResponseModel(entity *entities.Game) *models.GameResponse {
	createdAt, _ := time.Parse(time.RFC3339, entity.CreatedAt.Format(time.RFC3339))
	startAt, _ := time.Parse(time.RFC3339, entity.StartAt.Format(time.RFC3339))
	endAt, _ := time.Parse(time.RFC3339, entity.EndAt.Format(time.RFC3339))

	return &models.GameResponse{
		ID:          entity.ID.Hex(),
		Name:        entity.Name,
		Description: entity.Description,
		URLImage:    entity.URLImage,
		StartAt:     startAt,
		EndAt:       endAt,
		CreatedAt:   createdAt,
	}
}
