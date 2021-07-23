package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"valorize-app/config"
	"valorize-app/db"
	"valorize-app/handlers"
	appmiddleware "valorize-app/middleware"
	"valorize-app/services/ethereum"
)

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("username").(*jwt.Token)
	fmt.Print("\n\n" + string(user.Raw) + "\n\n")
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	cfg := config.NewConfig()
	dbInstance := db.Init(cfg)
	auth := handlers.AuthHandler{DB: dbInstance}
	ethInstance, err := ethereum.Connect()
	if err != nil {
		println("Error connecting to ethereum")
	}
	eth := handlers.EthHandler{
		Connection: ethInstance,
		DB:         dbInstance,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://valorize.local:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Static("/*", "app/dist")
	e.GET("/public", accessible)
	e.POST("/login", auth.Login)
	e.POST("/register", auth.Register)

	e.GET("/eth", eth.Ping)

	r := e.Group("/admin", appmiddleware.AuthMiddleware)
	r.POST("/wallet", eth.CreateWalletFromRequest)

	api := e.Group("/api/v0")
	api.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "All systems GO")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
