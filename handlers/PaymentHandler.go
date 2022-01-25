package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"valorize-app/contracts"
	"valorize-app/models"
	"valorize-app/services"
	"valorize-app/services/ethereum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/webhook"
)

type PaymentHandler struct {
	Server *Server
}

func NewPaymentHandler(s *Server) *PaymentHandler {
	return &PaymentHandler{s}
}

func (payment *PaymentHandler) CreateCheckoutSession(c echo.Context) error {
	stripe.Key = os.Getenv("STRIPE_KEY")
	user, _ := services.AuthUser(c, *payment.Server.DB)
	tokenName := c.FormValue("tokenName")
	tokenSymbol := c.FormValue("tokenSymbol")

	params := &stripe.CheckoutSessionParams{
		CustomerEmail: &user.Email,
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Deploy " + tokenName + " (" + tokenSymbol + ")"),
					},
					UnitAmount: stripe.Int64(1000),
				},
				Quantity: stripe.Int64(1),
			},
		},
		ClientReferenceID: stripe.String(strconv.FormatUint(uint64(user.ID), 10)),
		SuccessURL:        stripe.String(os.Getenv("FRONTEND_URL") + "/" + user.Username),
		CancelURL:         stripe.String(os.Getenv("FRONTEND_URL") + "/edit-profile"),
	}
	params.AddMetadata("name", tokenName)
	params.AddMetadata("symbol", tokenSymbol)

	fmt.Printf("name: %v, symbol: %v", tokenName, tokenSymbol)
	s, err := session.New(params)

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, s.URL)
}
func (payment *PaymentHandler) _fulfillOrder(session stripe.CheckoutSession) (common.Address, *types.Transaction, *contracts.CreatorToken, error) {
	client, err := ethereum.MainnetConnection()

	addr, tx, instance, err := ethereum.LaunchContract(client, session.Metadata["name"], session.Metadata["symbol"], "MAINNET")

	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	return addr, tx, instance, err
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
	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
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
		addr, tx, _, err := payment._fulfillOrder(session)
		customerId32bit, _ := strconv.ParseUint(session.ClientReferenceID, 10, 32)
		user, err := models.GetUserByID(uint(customerId32bit), *payment.Server.DB)

		fmt.Printf(`
==============================================================================================
    contract by owner %v deployed to address %v
==============================================================================================
`, user.ID, addr.String())
		if err != nil {
			payment.Server.Echo.Logger.Error(map[string]string{
				"message": "error accepting payment for " + string(session.ClientReferenceID),
			})
		}

		creatorToken := models.Token{
			UserId:          user.ID,
			ContractVersion: "v0.2.2",
			Name:            session.Metadata["name"],
			Symbol:          session.Metadata["symbol"],
			ChainId:         "1",
			OwnerAddress:    os.Getenv("HOTWALLET"),
			Address:         addr.String(),
			TxHash:          tx.Hash().String(),
		}

		err = payment.Server.DB.Create(&creatorToken).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "could not store contract information: " + err.Error(),
			})
		}

		user.HasDeployedToken = true
		err = payment.Server.DB.Save(&user).Error

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "could not update user info: " + err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"data": "payment accepted",
		})
	}
	return nil
}
