package api

import (
	"fmt"
	"howardsolutions/go/crypto/datatypes"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (datatypes.Rate, error) {
	upperCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))
}
