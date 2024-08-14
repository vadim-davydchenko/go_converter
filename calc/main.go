package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var result float64

	operation := getInputOperation()
	numbers := getInputNumber()

	switch operation {
	case "AVG":
		result = OperationAVG(numbers)
	case "SUM":
		result = OperationSUM(numbers)
	case "MED":
		result = OperationMED(numbers)
	}

	fmt.Printf("The Result of your %v operation = %v", operation, result)
}

func getInputOperation() string {
	var operation string
	for {
		fmt.Print("Select the operation to be calculated (AVG, SUM, MED):")
		fmt.Scan(&operation)
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			break
		}
		fmt.Println("Invalid operation. Please enter AVG, SUM, MED")
	}
	return operation
}

func getInputNumber() []float64 {
	var listNumbers string
	var sliceNumbers []float64

	fmt.Print("Input numbers for operation:")
	fmt.Scan(&listNumbers)

	stringNumbers := strings.Split(listNumbers, ",")

	for _, string := range stringNumbers {
		num, err := strconv.ParseFloat(string, 64)

		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			continue
		}

		sliceNumbers = append(sliceNumbers, num)
	}

	return sliceNumbers
}

func OperationAVG(numbers []float64) float64 {
	var sum float64
	lengthSlice := len(numbers)

	if lengthSlice == 0 {
		return 0
	}

	for _, num := range numbers {
		sum += num
	}
	AVG := sum / float64(lengthSlice)
	return AVG
}

func OperationSUM(numbers []float64) float64 {
	var SUM float64
	lengthSlice := len(numbers)

	if lengthSlice == 0 {
		return 0
	}

	for _, num := range numbers {
		SUM += num
	}
	return SUM
}

func OperationMED(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sort.Float64s(numbers)
	middleNumber := len(numbers) / 2

	if len(numbers)%2 != 0 {
		return numbers[middleNumber]
	}
	return (numbers[middleNumber-1] + numbers[middleNumber]) / 2
}
