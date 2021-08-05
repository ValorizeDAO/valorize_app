package main

import (
	"valorize-app/config"
	"valorize-app/handlers"
	"valorize-app/handlers/router"
)

func main() {
	cfg := config.NewConfig()
	s := handlers.NewServer(cfg)
	e := router.NewRouter(s)
	e.Logger.Fatal(e.Start(":1323"))
}
