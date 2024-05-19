package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Структура для парсингу JSON відповіді
type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}

func UsdRate() (float64, error) {
	url := "https://open.er-api.com/v6/latest/USD"

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var er ExchangeRates
	if err := json.Unmarshal(body, &er); err != nil {
		return 0, err
	}

	rate, ok := er.Rates["UAH"]
	if !ok {
		return 0, fmt.Errorf("could not find UAH rate in response")
	}

	return rate, nil
}
