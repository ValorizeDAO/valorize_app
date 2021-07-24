package handlers

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"valorize-app/config"
	"valorize-app/db"
	"valorize-app/services/ethereum"
)

type Server struct {
	Echo       *echo.Echo
	DB         *gorm.DB
	Config     *config.Config
	BlockChain *ethclient.Client
}

func NewServer(cfg *config.Config) *Server {
	ethInstance, err := ethereum.Connect()
	if err != nil {
		fmt.Println("Error connecting to Ethereum")
	}
	return &Server{
		Echo:       echo.New(),
		DB:         db.Init(cfg),
		Config:     cfg,
		BlockChain: ethInstance,
	}
}
