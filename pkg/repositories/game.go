package repositories

import (
	"context"

	"github.com/getmilly/games/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//GameRepository ...
type GameRepository interface {
	Insert(*entities.Game) error
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
func (repository gameRepository) Insert(game *entities.Game) error {
	res, err := repository.collection.InsertOne(context.Background(), game)
	game.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

//FindByID
func (repository gameRepository) FindByID(ID string) (*entities.Game, error) {
	ctx := context.Background()

	game := new(entities.Game)

	objectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	err = repository.collection.FindOne(ctx, bson.M{
		"_id": objectID,
	}).Decode(game)

	if err != nil {
		return nil, err
	}

	return game, nil
}
