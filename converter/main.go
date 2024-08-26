package main

import (
	"fmt"
)

const (
	USDToEUR = 0.92
	USDToRUB = 85.41
	EURToUSD = 1 / USDToEUR
	EURToRUB = USDToRUB / USDToEUR
	RUBToUSD = 1 / USDToRUB
	RUBToEUR = 1 / EURToRUB
)

var currencyMap = map[string]float64{
	"USDEUR": USDToEUR,
	"USDRUB": USDToRUB,
	"EURUSD": EURToUSD,
	"EURRUB": EURToRUB,
	"RUBUSD": RUBToUSD,
	"RUBEUR": RUBToEUR,
}

func main() {
	fmt.Println("__Converter currency__")
	sourceCurrency, amount, targetCurrency := getInput()
	convertAmount := convertCurrency(&currencyMap, sourceCurrency, amount, targetCurrency)
	fmt.Printf("Converted amount from %v %v = %.2f %v\n", amount, sourceCurrency, convertAmount, targetCurrency)
}

func getInput() (string, float64, string) {
	var sourceCurrency string
	var amount float64
	var targetCurrency string
	for {
		fmt.Print("Entering the source currency (USD,EUR,RUB):")
		fmt.Scan(&sourceCurrency)
		if sourceCurrency == "USD" || sourceCurrency == "EUR" || sourceCurrency == "RUB" {
			break
		}
		fmt.Println("Invalid source currency. Please enter USD,EUR,RUB")
	}

	for {
		fmt.Print("Entering the amount:")
		_, err := fmt.Scan(&amount)
		if amount > 0 && err == nil {
			break
		}
		fmt.Println("Invalid amount.Please enter a number and the number must be greater than zero")
		fmt.Scanln()
	}

	for {
		switch sourceCurrency {
		case "USD":
			fmt.Print("Entering the target currency (EUR,RUB):")
		case "EUR":
			fmt.Print("Entering the target currency (USD,RUB):")
		case "RUB":
			fmt.Print("Entering the target currency (USD,EUR):")
		}
		fmt.Scan(&targetCurrency)
		if sourceCurrency == "USD" && (targetCurrency == "EUR" || targetCurrency == "RUB") ||
			(sourceCurrency == "EUR" && (targetCurrency == "USD" || targetCurrency == "RUB")) ||
			(sourceCurrency == "RUB" && (targetCurrency == "USD" || targetCurrency == "EUR")) {
			break
		}
		fmt.Println("Invalid target currency. Please enter a valid currency according to the prompts")
	}
	return sourceCurrency, amount, targetCurrency
}

func convertCurrency(currencyMap *map[string]float64, sourceCurrency string, amount float64, targetCurrency string) float64 {
	currencynKey := sourceCurrency + targetCurrency
	if course, exists := (*currencyMap)[currencynKey]; exists {
		return amount * course
	}
	return 0
}
