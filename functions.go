package main

import fmt "FMT"

func sum(number1 int, number2 int) int {
	return number1 + number2
}

func multiply(number1, number2 int) int {
	return number1 * number2
}

func main() {
	number1 := 3
	number2 := 2

	res1 := sum(number1, number2)
	fmt.Printf("The sum is %v\n", res1)

	res2 := multiply(number1, number2)
	fmt.Printf("The product is %v\n", res2)
}
