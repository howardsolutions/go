package main

import "howardsolutions/go/crypto/api"

func main() {
	rate, err := api.GetRate("BTC")
	print(rate, err)
}
