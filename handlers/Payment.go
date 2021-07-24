package handlers

import (
  "encoding/json"
  "fmt"
  "github.com/labstack/echo/v4"
  "github.com/stripe/stripe-go/v72"
  "github.com/stripe/stripe-go/v72/checkout/session"
  "github.com/stripe/stripe-go/v72/webhook"
  "io/ioutil"
  "net/http"
  "os"
)

type PaymentHandler struct{
  Server *Server
}

func NewPaymentHandler(s *Server) *PaymentHandler {
  return &PaymentHandler{s }
}

func (payment *PaymentHandler) CreateCheckoutSession(c echo.Context) (err error) {
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
func fulfillOrder(session stripe.CheckoutSession) {
  fmt.Println(session.Customer)
  e, err := json.Marshal(session.Customer)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(e))
  fmt.Println("\n\n==============\n\n    TADA    \n\n============")
}

func (payment *PaymentHandler) OnPaymentAccepted(c echo.Context) error {
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
    fulfillOrder(session)
  }

  return c.JSON(http.StatusOK, map[string]string{
    "data": "payment accepted",
  })
}
