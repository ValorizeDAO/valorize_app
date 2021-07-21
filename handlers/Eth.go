package handlers

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EthHandler struct {
	Connection *ethclient.Client
}

func (eth *EthHandler) Ping(c echo.Context) error {
	connection, err := eth.Connection.NetworkID(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not connect to Ethereum Blockchain",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"data": "connected to " + connection.String(),
	})
}
