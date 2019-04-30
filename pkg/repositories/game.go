package repositories

import (
	"context"

	"github.com/getmilly/games/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//GameRepository ...
type GameRepository interface {
	Insert(entities.Game) error
	FindByID(ID string) (*entities.Game, error)
}

//gameRepository ...
type gameRepository struct {
	collection *mongo.Collection
}

//NewGameRepository ...
func NewGameRepository(client *mongo.Client) GameRepository {
	return &gameRepository{
		collection: client.Database("games").Collection("games"),
	}
}

//Insert ...
func (repository gameRepository) Insert(game entities.Game) error {
	_, err := repository.collection.InsertOne(context.Background(), game)
	return err
}

//FindByID
func (repository gameRepository) FindByID(ID string) (*entities.Game, error) {
	ctx := context.Background()

	game := new(entities.Game)

	result := repository.collection.FindOne(ctx, bson.M{
		"game_id": ID,
	})

	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Decode(game)

	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	return game, nil
}
