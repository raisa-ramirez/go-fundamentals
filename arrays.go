package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 10
	fmt.Println(a)
	fmt.Println(a[0])

	var people [4]string = [4]string{"Juan", "Marcos", "Sara", "Elías"}
	for i := range 4 {
		fmt.Println(people[i])
	}

	animals := [2][2]string{{"BAO", "PELUSA"}, {"KIARA", "NALA"}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%v ", animals[i][j])
		}
		fmt.Println("")
	}

	var initialArray [5]int
	i := 1
	for i < 6 {
		initialArray[i-1] = i * 2
		i++
	}

	for j := range 5 {
		fmt.Println(initialArray[j])
	}

	fmt.Printf("Array size %v", len(initialArray))
}
