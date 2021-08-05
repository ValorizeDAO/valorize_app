package main

import (
	"os"
	"valorize-app/config"
	"valorize-app/handlers"
	"valorize-app/handlers/router"
)

func main() {
	cfg := config.NewConfig()
	s := handlers.NewServer(cfg)
	e := router.NewRouter(s)
	if os.Getenv("ENVIRONMENT") == "DEVELOPMENT" {
		e.Logger.Fatal(e.Start(":1323"))
	} else {
		e.Logger.Fatal(e.Start(":8080"))
	}
}
