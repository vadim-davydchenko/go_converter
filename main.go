package main

import "fmt"

const USDToEUR = 0.92
const USDToRUB = 85.41

func main() {
	fmt.Println("__Converter currency__")
	sourceCurrency, amount, targetCurrency := getInput()
	fmt.Println(sourceCurrency)
	fmt.Println(amount)
	fmt.Println(targetCurrency)
	// EURToRUB := USDToRUB / USDToEUR
	// fmt.Println(EURToRUB)

}

func getInput() (string, float64, string) {
	var sourceCurrency string
	var amount float64
	var targetCurrency string
	fmt.Print("Entering the source currency (USD,EUR,RUB):")
	fmt.Scan(&sourceCurrency)
	fmt.Print("Entering the amount:")
	fmt.Scan(&amount)
	switch sourceCurrency {
	case "USD":
		fmt.Print("Entering the target currency (EUR,RUB):")
	case "EUR":
		fmt.Print("Entering the target currency (USD,RUB):")
	case "RUB":
		fmt.Print("Entering the target currency (USD,EUR):")
	}
	fmt.Scan(&targetCurrency)
	return sourceCurrency, amount, targetCurrency
}

// func checkCurrencyAmount() bool {
// 	var userChoise string
// }

// func convertCurrency(amount float64, sourceCurrency string, targetCurrency string) float64 {
// 	return 0
// }
