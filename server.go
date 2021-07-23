package main

import (
  "encoding/json"
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/stripe/stripe-go/v72"
  "github.com/stripe/stripe-go/v72/checkout/session"
  "github.com/stripe/stripe-go/v72/webhook"
  "io/ioutil"
  "net/http"
  "os"
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
  ethInstance, err := ethereum.Connect()
  stripe.Key = "sk_test_51JGBbjBhSkl0qU1AdCzBjVv6N0Z2xyYqHTfYPOECkuFdl4lA9fyLIz6lHrKP702wlybuwcfh1rB7ljG8zUzzta7k00ytyRYt2d"
  if err != nil {
    println("Error connecting to ethereum")
  }

  auth := handlers.AuthHandler{DB: dbInstance}
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

  //e.Static("/*", "app/dist")
  e.GET("/public", accessible)
  e.GET("/success", func(c echo.Context) error {
    return c.String(http.StatusOK, "Success")
  })
  e.GET("/cancel", func(c echo.Context) error {
    return c.String(http.StatusOK, "Payment error")
  })
  e.POST("/login", auth.Login)
  e.POST("/register", auth.Register)
  e.POST("/create-checkout-session", createCheckoutSession)
  e.GET("/eth", eth.Ping)
  e.POST("/payments/successhook", onPaymentAccepted)

  r := e.Group("/admin", appmiddleware.AuthMiddleware)
  r.POST("/wallet", eth.CreateWalletFromRequest)

  api := e.Group("/api/v0")
  api.GET("/healthcheck", func(c echo.Context) error {
    return c.String(http.StatusOK, "All systems GO")
  })
  e.Logger.Fatal(e.Start(":1323"))
}

func createCheckoutSession(c echo.Context) (err error) {
  params := &stripe.CheckoutSessionParams{
    PaymentMethodTypes: stripe.StringSlice([]string{
      "card",
    }),
    Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
    LineItems: []*stripe.CheckoutSessionLineItemParams{
      {
        PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
          Currency: stripe.String("usd"),
          ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
            Name: stripe.String("Deploy Token"),
          },
          UnitAmount: stripe.Int64(1000),
        },
        Quantity: stripe.Int64(1),
      },
    },
    //TODO: use environment variable
    SuccessURL: stripe.String("http://localhost:1323/success"),
    CancelURL:  stripe.String("http://localhost:1323/cancel"),
    //TODO: have to connect to the user and pass local user data
    //ClientReferenceID: services.AuthUser(c, checkout.DB).ID
    //user, _ := services.AuthUser(c, db)
    //user.ID
  }

  s, _ := session.New(params)

  if err != nil {
    return err
  }

  return c.Redirect(http.StatusSeeOther, s.URL)
}
func FulfillOrder(session stripe.CheckoutSession) {
  fmt.Println(session.Customer)
  e, err := json.Marshal(session.Customer)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(e))
  fmt.Println("\n\n==============\n\n    TADA    \n\n============")
}

func onPaymentAccepted(c echo.Context) error {
  const MaxBodyBytes = int64(65536)
  requestBody := http.MaxBytesReader(c.Response().Writer, c.Request().Body, MaxBodyBytes)

  body, err := ioutil.ReadAll(requestBody)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
    return c.JSON(http.StatusInternalServerError, map[string]string{
      "error": "could not process payment",
    })
  }

  // Pass the request body and Stripe-Signature header to ConstructEvent, along with the webhook signing key
  // You can find your endpoint's secret in your webhook settings
  endpointSecret := "whsec_rJO3g8YgfWas3uDM4axkZ1Dj1bC2xwnU"
  event, err := webhook.ConstructEvent(body, c.Request().Header.Get("Stripe-Signature"), endpointSecret)

  if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{
      "error": "could not authenticate signature",
    })
  }

  // Handle the checkout.session.completed event
  if event.Type == "checkout.session.completed" {
    var session stripe.CheckoutSession
    err := json.Unmarshal(event.Data.Raw, &session)
    if err != nil {
      return c.JSON(http.StatusInternalServerError, map[string]string{
        "error": fmt.Sprintf("Error parsing webhook JSON: %v\n", err),
      })
    }
    FulfillOrder(session)
  }

  return c.JSON(http.StatusOK, map[string]string{
    "data": "payment accepted",
  })
}
