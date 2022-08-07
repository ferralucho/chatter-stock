package services

import (
	"github.com/ferralucho/chatter-stock/src/domain/bots"
	"github.com/ferralucho/chatter-stock/src/domain/stocks"
	"github.com/ferralucho/chatter-stock/src/repository/rest"
	"github.com/ferralucho/chatter-stock/src/rest_errors"
	"github.com/ferralucho/chatter-stock/src/utils/date_utils"
)

var (
	BotsService botsServiceInterface = &botsService{}
)

type botsService struct{}

type botsServiceInterface interface {
	CreateCommand(command bots.Command) (*stocks.StockData, rest_errors.RestErr)
}

func (s *botsService) CreateCommand(command bots.Command) (*stocks.StockData, rest_errors.RestErr) {
	if err := command.Validate(); err != nil {
		return nil, err
	}
	command.DateCreated = date_utils.GetNowDBFormat()
	botsRepository := rest.NewRestStocksRepository()
	var stock *stocks.StockData
	var err error

	if command.CommandType == bots.CommandTypeStock {
		stock, err = botsRepository.GetStock(command)
		if err != nil {
			return nil, rest_errors.NewBadRequestError("error getting stock")
		}
	}

	return stock, nil
}
