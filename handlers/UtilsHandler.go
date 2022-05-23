package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"valorize-app/models"

	"github.com/labstack/echo/v4"
)

type UtilsHandler struct {
	server *Server
	models *models.Model
}

func NewUtilsHandler(s *Server, m *models.Model) *UtilsHandler {
	return &UtilsHandler{s, m}
}

func (utils *UtilsHandler) ShowEthPrice(c echo.Context) error {
	url := "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD&api_key=" + os.Getenv("CRYPTOCOMPARE_KEY")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, json.RawMessage(body))
}
