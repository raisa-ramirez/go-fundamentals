package main

import "fmt"

func sumValues(numbers ...int) {
	total := 0

	for i := range numbers {
		total += numbers[i]
	}
	// for _, v := range numbers {
	// 	total += v
	// }
	fmt.Println(total)
}

func main() {
	sumValues(3, 3, 3)
	sumValues(9, 9, 9)
	data := []int{1, 2, 3, 4, 5, 6}
	sumValues(data...)
}
