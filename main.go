package main

import "fmt"

func main() {
	const USDToEUR = 0.92
	const USDToRUB = 85.41
	EURToRUB := USDToRUB / USDToEUR
	fmt.Println(EURToRUB)
}
