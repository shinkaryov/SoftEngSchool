package exchange_rate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ExchangeRate represents the exchange rate data
type ExchangeRate struct {
	R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float64 `json:"rate"`
	Cc           string  `json:"cc"`
	Exchangedate string  `json:"exchangedate"`
}

// GetRateFromUrl fetches the exchange rate from the given URL
func GetRateFromUrl(url string) (float64, error) {
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Cookie", "__cf_bm=5gF6gAzSNvHz.FookaPaHMHnfIA6PYVvl10BidGBGvU-1716056860-1.0.1.1-b2vt11TqvuFG.aAqU0ThQvgkO0BF3yyLaQbsVuWdLh0VGMVqJ0XyREnRf8Jemq0.xVu0EJGFEUfGtmmnXo16Y4pD8rMLaTmK.1YcGT_1CCI")

	res, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to perform request: %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}

	var exchangeRates []ExchangeRate
	err = json.Unmarshal(body, &exchangeRates)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	if len(exchangeRates) == 0 {
		return 0, fmt.Errorf("no exchange rates found in response")
	}

	return exchangeRates[0].Rate, nil
}
