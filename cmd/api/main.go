package main

import (
	"log"

	appctx "github.com/heru-oktafian/cms-be/internal/app"
	httpapp "github.com/heru-oktafian/cms-be/internal/delivery/http"
)

func main() {
	container := appctx.Bootstrap()
	app := httpapp.NewApp(container)
	log.Printf("%s running on %s:%s", container.Config.AppName, container.Config.AppHost, container.Config.AppPort)
	log.Fatal(app.Listen(container.Config.AppHost + ":" + container.Config.AppPort))
}
