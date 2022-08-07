package rest

import (
	"encoding/csv"
	"github.com/ferralucho/chatter-stock/src/domain/bots"
	"github.com/ferralucho/chatter-stock/src/domain/stocks"
	"github.com/ferralucho/chatter-stock/src/rest_errors"
	"io"
	"net/http"
	"net/url"
)

var (
	BaseURL = "https://stooq.com/q/l/?f=sd2t2ohlcv&h&e=csv&s="
)

type RestBotsRepository interface {
	GetStock(command bots.Command) (*stocks.StockData, rest_errors.RestErr)
}

type stocksRepository struct{}

func NewRestStocksRepository() RestBotsRepository {
	return &stocksRepository{}
}

func (r *stocksRepository) GetStock(command bots.Command) (*stocks.StockData, rest_errors.RestErr) {
	params := url.Values{}
	params.Add("s", command.Message)

	resp, err := http.Get(BaseURL + params.Encode())
	if resp == nil || err != nil {
		return nil, rest_errors.NewInternalServerError("invalid response when trying to get stock", nil)
	}

	results, err := ReadCSVFromHttpRequest(resp)
	if results == nil || err != nil {
		return nil, rest_errors.NewInternalServerError("csv parse error", nil)
	}
	var stock stocks.StockData

	stock = stocks.StockData{
		Symbol: results[0][0],
	}

	return &stock, nil
}

func ReadCSVFromHttpRequest(resp *http.Response) ([][]string, error) {
	reader := csv.NewReader(resp.Body)
	var results [][]string
	for {
		// read one row from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		results = append(results, record)
	}
	return results, nil
}
