package main

import (
	"github.com/getmilly/games/api/controllers"
	"github.com/getmilly/grok/api"
	gnats "github.com/getmilly/grok/nats"
	"github.com/sarulabs/di"
)

func main() {
	settings := api.SettingsFromDotEnv()

	server := api.ConfigureServer(settings, api.DefaultHealthChecks())

	server.AddController(di.Def{
		Name:  "app-controller",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			producer := c.Get("nats-producer").(*gnats.Producer)
			return controllers.NewAppController(producer), nil
		},
	})

	server.AddDependency(di.Def{
		Name:  "nats-producer",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			conn := c.Get("nats-conn").(*nats.EncodedConn)
			return gnats.NewProducer(conn), nil
		},
	})

	server.AddDependency(di.Def{
		Name:  "nats-conn",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			return gnats.Connect(nats.DefaultURL, "")
		},
	})

	server.Run()
}
