package main

import (
	"github.com/caarlos0/env"
	"github.com/getmilly/games/api/controllers"
	"github.com/getmilly/games/api/settings"
	"github.com/getmilly/games/pkg/repositories"
	"github.com/getmilly/games/pkg/services"
	"github.com/getmilly/grok/api"
	"github.com/getmilly/grok/mongodb"
	"github.com/joho/godotenv"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	baseSettings := api.SettingsFromDotEnv()
	server := api.ConfigureServer(baseSettings, api.DefaultHealthChecks())

	appSettings := readSettings()

	server.AddController(di.Def{
		Name:  "game-controller",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			gameService := c.Get("game-service").(services.GameService)
			return controllers.NewGameController(gameService), nil
		},
	})

	server.AddDependency(di.Def{
		Name:  "mongo-client",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			return mongodb.Connect(appSettings.MongoConnectionString)
		},
	})

	server.AddDependency(di.Def{
		Name:  "game-service",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			gameRepository := c.Get("game-repository").(repositories.GameRepository)
			return services.NewGameService(gameRepository), nil
		},
	})

	server.AddDependency(di.Def{
		Name:  "game-repository",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			client := c.Get("mongo-client").(*mongo.Client)
			return repositories.NewGameRepository(client), nil
		},
	})

	server.Run()
}

func readSettings() *settings.Settings {
	settings := new(settings.Settings)

	godotenv.Load()
	err := env.Parse(settings)

	if err != nil {
		panic(err)
	}

	return settings
}
