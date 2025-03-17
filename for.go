package main

import "fmt"

func main() {
	const format = "%v x %v = %v\n"

	for i := range 11 {
		table := 2
		total := table * i
		fmt.Printf(format, table, i, total)
	}

	fmt.Println("")

	for i := 1; i < 11; i++ {
		table := 3
		total := table * i
		fmt.Printf(format, table, i, total)
	}

	fmt.Println("")
	i := 1
	for i < 11 {
		table := 4
		total := table * i
		fmt.Printf(format, table, i, total)
		i++
	}
}
