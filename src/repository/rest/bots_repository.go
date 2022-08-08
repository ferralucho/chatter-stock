package rest

import (
	"encoding/csv"
	"github.com/ferralucho/chatter-stock/src/domain/bots"
	"github.com/ferralucho/chatter-stock/src/domain/stocks"
	"github.com/ferralucho/chatter-stock/src/rest_errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
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

	open, high, low, close, volume := parseStockResults(err, results)

	stock = stocks.StockData{
		Symbol: results[1][0],
		Date:   results[1][1],
		Time:   results[1][2],
		Open:   float32(open),
		High:   float32(high),
		Low:    float32(low),
		Close:  float32(close),
		Volume: int32(volume),
	}

	return &stock, nil
}

func parseStockResults(err error, results [][]string) (float64, float64, float64, float64, float64) {
	open, err := strconv.ParseFloat(results[1][3], 32)
	if err != nil {
		open = 0.0
	}

	high, err := strconv.ParseFloat(results[1][4], 32)
	if err != nil {
		open = 0.0
	}
	low, err := strconv.ParseFloat(results[1][5], 32)
	if err != nil {
		open = 0.0
	}
	close, err := strconv.ParseFloat(results[1][6], 32)
	if err != nil {
		open = 0.0
	}

	volume, err := strconv.ParseFloat(results[1][7], 32)
	if err != nil {
		open = 0.0
	}
	return open, high, low, close, volume
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
