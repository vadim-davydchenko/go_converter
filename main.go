package main

import "fmt"

func readInputCurrency() float64 {
	var amount float64
	fmt.Print("Enter the amount of the original currency:")
	fmt.Scan(&amount)
	return amount
}

func convertCurrency(amount float64, sourceCurrency string, targetCurrency string) float64 {

}

func main() {
	const USDToEUR = 0.92
	const USDToRUB = 85.41

	// userInput := readInputCurrency()
	EURToRUB := USDToRUB / USDToEUR
	fmt.Println(EURToRUB)
}
