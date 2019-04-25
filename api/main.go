package main

import (
	"github.com/getmilly/games/api/controllers"
	"github.com/getmilly/grok/api"
	"github.com/sarulabs/di"
)

func main() {
	settings := api.SettingsFromDotEnv()

	server := api.ConfigureServer(settings, api.DefaultHealthChecks())

	server.AddController(di.Def{
		Name:  "games-controller",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			return controllers.NewGamesController(), nil
		},
	})

	server.Run()
}
