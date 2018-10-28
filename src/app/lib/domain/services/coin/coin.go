package service

import (
	config "app/assets/config/api/coin-market-cap"
	"net/http"
	"strings"
)

type Coin struct{}

func (c Coin) GetTickerBySymbols(symbols []string) *http.Response {
	query := strings.Join(symbols, ",")
	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=" + query
	req, err := http.NewRequest("GET", url, nil)
	key := config.CoinMarketCap()
	req.Header.Set("X-CMC_PRO_API_KEY", key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp
}
