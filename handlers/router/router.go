package router

import (
	"net/http"
	"os"
	"valorize-app/handlers"
	appmiddleware "valorize-app/middleware"
	"valorize-app/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(s *handlers.Server) echo.Echo {
	e := *s.Echo
	m := models.NewModels(s.DB)
	payment := handlers.NewPaymentHandler(s, m)
	auth := handlers.NewAuthHandler(s, m)
	eth := handlers.NewEthHandler(s, m)
	user := handlers.NewUserHandler(s, m)
	token := handlers.NewTokenHandler(s, m)
	wallet := handlers.NewWalletHandler(s, m)
	utils := handlers.NewUtilsHandler(s, m)
	contracts := handlers.NewContractsHandler(s, m)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("ALLOW_ORIGIN")},
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
	api.GET("/get-gas-price", token.GetGasPriceToLaunchToken)

	tokenPublic := api.Group("/token")
	tokenPublic.GET("/:id", token.ShowToken)
	tokenPublic.GET("/:id/admins", token.ShowTokenAdmins)
	tokenPublic.GET("/:id/airdrops/:address/claim", token.AirdropClaimAmount)
	tokenPublic.GET("/:address/airdrops", token.AirdropEligibility)

	me := api.Group("/me", appmiddleware.AuthMiddleware)
	me.GET("", auth.ShowUser)
	me.PUT("/picture", auth.UpdatePicture)
	me.PUT("/profile", auth.UpdateProfile)
	me.PUT("/links", auth.UpdateLinks)
	me.DELETE("/links", auth.DeleteLinks)
	me.PUT("/tokendata", auth.UpdateTokenData)
	me.GET("/wallets", wallet.Index)
	myTokens := me.Group("/tokens")
	myTokens.GET("", token.Index)
	myTokens.GET("/:id/balance", token.GetCoinBalanceForAuthUser)
	myTokens.PUT("/:id/airdrop/create", token.NewAirdrop)
	me.GET("/contract-bytecode/:key", contracts.GetContractBytecode)

	r := api.Group("/admin", appmiddleware.AuthMiddleware)
	r.POST("/wallet", eth.CreateWalletFromRequest)
	r.POST("/wallet/new", eth.AddWalletToAccount)
	r.POST("/deploy", eth.DeployCreatorToken)
	r.POST("/deploy/simpletoken", eth.DeploySimpleToken)

	userGroup := api.Group("/users")
	userGroup.GET("/:username", user.Show)
	userGroup.GET("/:username/links", user.GetLinks)
	userGroup.GET("/:username/token", token.Show)
	userGroup.POST("/:username/token/stakingrewards", token.GetTokenStakingRewards)
	userGroup.POST("/:username/token/sellingrewards", token.GetTokenSellingRewards)

	u := api.Group("/utils")
	u.GET("/price", utils.ShowEthPrice)

	return e
}
