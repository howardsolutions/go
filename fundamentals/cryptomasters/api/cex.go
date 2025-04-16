package api

import (
	"fmt"
	"howardsolutions/go/crypto/datatypes"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	upperCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		json := string(bodyBytes)
		fmt.Println(json)
	} else {
		return nil, fmt.Errorf("Status code received: %v", res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: 20}

	return &rate, nil
}
