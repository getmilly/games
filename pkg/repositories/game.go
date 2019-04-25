package repositories

import "go.mongodb.org/mongo-driver/mongo"

//GameRepository ...
type GameRepository interface{}

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
