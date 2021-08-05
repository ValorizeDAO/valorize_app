package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"os"
)

type UtilsHandler struct {
	server *Server
}

func NewUtilsHandler(s *Server) *UtilsHandler {
	return &UtilsHandler{s}
}

func (utils *UtilsHandler) ShowEthPrice(c echo.Context) error {
  fmt.Println("TEST")
	url := "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD&api_key=" + os.Getenv("CRYPTOCOMPARE_KEY")
	method := "GET"

	client := &http.Client {
	}
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
	fmt.Println(string(body))
	//response, err := json.Marshal(body)
	return c.JSON(http.StatusOK, json.RawMessage(body))
}
