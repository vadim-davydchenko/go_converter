package main

import (
	"fmt"
)

func main() {
	operation := getInputOperation()
	fmt.Println(operation)
}

func getInputOperation() string {
	var operation string
	for {
		fmt.Print("Select the operation to be calculated (AVG,SUM,MED):")
		fmt.Scan(&operation)
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			break
		}
		fmt.Println("Invalid operation. Please enter AVG,SUM,MED")
	}
	return operation
}
