package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/webhook"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"valorize-app/contracts"
	"valorize-app/models"
	"valorize-app/services"
	"valorize-app/services/ethereum"
)

type PaymentHandler struct {
	Server *Server
}

func NewPaymentHandler(s *Server) *PaymentHandler {
	return &PaymentHandler{s}
}

func (payment *PaymentHandler) CreateCheckoutSession(c echo.Context) error {
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
						Name: stripe.String("Deploy Token"),
					},
					UnitAmount: stripe.Int64(1000),
				},
				Quantity: stripe.Int64(1),
			},
		},
		ClientReferenceID: stripe.String(strconv.FormatUint(uint64(user.ID), 10)),
		//TODO: use environment variable
		SuccessURL: stripe.String(os.Getenv("FRONTEND_URL") + "/" + user.Username),
		CancelURL:  stripe.String(os.Getenv("FRONTEND_URL") + "/edit-profile"),
	}
	params.AddMetadata("token", tokenName)
	params.AddMetadata("symbol", tokenSymbol)

	s, err := session.New(params)

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, s.URL)
}
func (payment *PaymentHandler) _fulfillOrder(session stripe.CheckoutSession) (common.Address, *types.Transaction, *contracts.CreatorToken, error){
	var clientUrl string
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		clientUrl = os.Getenv("MAINNET_NODE")
	} else {
		clientUrl = os.Getenv("ETH_NODE")
	}
	client, err := ethclient.Dial(clientUrl)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("=======================\n\nConnected to ethereum mainnet\n\n=======================")
	}

	addr, tx, instance, err := ethereum.LaunchContract(client, session.Metadata["name"], session.Metadata["symbol"])

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
		addr, tx, _, err := payment._fulfillOrder(session)
		customerId32bit, _ := strconv.ParseUint(session.ClientReferenceID, 10, 32)
		user, err := models.GetUserByID(uint(customerId32bit), *payment.Server.DB)
		if err != nil {
			payment.Server.Echo.Logger.Error(map[string]string{
				"message": "error accepting payment for " + string(session.ClientReferenceID),
			})
		}

		creatorToken := models.Token{
			UserId:          user.ID,
			ContractVersion: "v0.0.1",
			Name:            session.Metadata["name"],
			Symbol:          session.Metadata["symbol"],
			Network:         "MAINNET",
			OwnerAddress:    os.Getenv("HOTWALLET"),
			Address:         addr.String(),
			User:            user,
			TxHash:          tx.Hash().String(),
		}

		err = payment.Server.DB.Create(&creatorToken).Error

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "could not store contract information: " + err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"data": "payment accepted",
		})
	}
	return nil
}
