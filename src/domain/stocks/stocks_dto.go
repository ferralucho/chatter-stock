package stocks

type StocksResponse struct {
	Symbols []StockData `json:"symbols"`
}

type StockData struct {
	Symbol string  `json:"id"`
	Date   string  `json:"date"`
	Time   string  `json:"time"`
	Open   float32 `json:"open"`
	High   float32 `json:"high"`
	Low    float32 `json:"low"`
	Close  float32 `json:"close"`
	volume int32   `json:"volume"`
}
