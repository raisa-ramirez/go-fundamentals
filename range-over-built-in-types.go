package main

import "fmt"

func main() {
	prices := []int{100, 200, 300}
	total := 0
	for _, price := range prices {
		total += price
	}

	fmt.Println("Total $", total)

	population := map[string]int{"women": 100, "men": 200}
	for i := range population {
		fmt.Printf("%s : %v\n", i, population[i])
	}
	fmt.Println("\n")
	for key, data := range population {
		fmt.Printf("%s : %v\n", key, data)
	}
}
