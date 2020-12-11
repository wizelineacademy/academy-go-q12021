package helpers

import (
	"encoding/json"
	"net/http"
	"strings"

	"../db/models"
)

func DecodeBody(req *http.Request) (models.Bitcoin, bool) {
	var btc models.Bitcoin
	err := json.NewDecoder(req.Body).Decode(&btc)
	if err != nil {
		return models.Bitcoin{}, false
	}

	return btc, true
}

func IsValidBase(base string) bool {
	b := strings.TrimSpace(base)
	if b != "BTC" {
		return false
	}

	return true
}

func IsValidCurrency(currency string) bool {
	c := strings.TrimSpace(currency)
	if c != "MXN" {
		return false
	}

	return true
}

func IsValidAmount(amount int) bool {
	if amount < 0 {
		return false
	}

	return true
}
