package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stripe/stripe-go/v72"
	"net/http"
	"os"
	"valorize-app/handlers"
	appmiddleware "valorize-app/middleware"
)

func NewRouter(s *handlers.Server) echo.Echo {
	stripe.Key = os.Getenv("STRIPE_KEY")
	e := *s.Echo
	payment := handlers.NewPaymentHandler(s)
	auth := handlers.NewAuthHandler(s)
	eth := handlers.NewEthHandler(s)
	user := handlers.NewUserHandler(s)
	wallet := handlers.NewWalletHandler(s)
	token := handlers.NewTokenHandler(s)
	utils := handlers.NewUtilsHandler(s)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	e.Static("/*", "app/dist")
	e.Static("/static/images", "dist/images")

	e.GET("/success", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success")
	})
	e.GET("/cancel", func(c echo.Context) error {
		return c.String(http.StatusOK, "Payment error")
	})
	e.POST("/login", auth.Login)
	e.GET("/logout", auth.Logout)
	e.POST("/register", auth.Register)
	e.GET("/create-checkout-session", payment.CreateCheckoutSession)
	e.GET("/eth", eth.Ping)
	e.POST("/payments/successhook", payment.OnPaymentAccepted)

	api := e.Group("/api/v0")

	me := api.Group("/me", appmiddleware.AuthMiddleware)
	me.GET("", auth.ShowUser)
	me.PUT("/picture", auth.UpdatePicture)
	me.PUT("/profile", auth.UpdateProfile)

	r := api.Group("/admin", appmiddleware.AuthMiddleware)
	r.POST("/wallet", eth.CreateWalletFromRequest)
	r.POST("/deploy", eth.DeployCreatorToken)


	userGroup := api.Group("/users")
	userGroup.GET("/:username", user.Show)
	userGroup.GET("/:username/token", token.Show)
	userGroup.GET("/:username/wallets", wallet.Index)

	u := api.Group("/utils")
	u.GET("/price", utils.ShowEthPrice)

	return e
}
