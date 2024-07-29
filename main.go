package main

import "fmt"

const USDToEUR = 0.92
const USDToRUB = 85.41

func getInput() float64 {
	var amount float64
	fmt.Print("Enter the amount of the original currency:")
	fmt.Scan(&amount)
	return amount
}

func convertCurrency(amount float64, sourceCurrency string, targetCurrency string) float64 {
	return 0
}

func main() {
	amount := getInput()

	EURToRUB := USDToRUB / USDToEUR
	fmt.Println(EURToRUB)

}
