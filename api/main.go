package main

import (
	"github.com/caarlos0/env"
	"github.com/getmilly/games/api/controllers"
	"github.com/getmilly/games/pkg/repositories"
	"github.com/getmilly/games/pkg/services"
	"github.com/getmilly/grok/api"
	"github.com/getmilly/grok/mongodb"
	"github.com/joho/godotenv"
	"github.com/sarulabs/di"
)

func main() {
	settings := api.SettingsFromDotEnv()
	server := api.ConfigureServer(settings, api.DefaultHealthChecks())

	appSettings := readSettings()

	client, err := mongodb.Connect(appSettings.MongoConnectionString)

	if err != nil {
		panic(err)
	}

	server.AddController(di.Def{
		Name:  "game-controller",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			gameService := c.Get("game-service").(services.GameService)
			return controllers.NewGameController(gameService), nil
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
			return repositories.NewGameRepository(client), nil
		},
	})

	server.Run()
}

func readSettings() *Settings {
	settings := new(Settings)

	godotenv.Load()
	err := env.Parse(settings)

	if err != nil {
		panic(err)
	}

	return settings
}
