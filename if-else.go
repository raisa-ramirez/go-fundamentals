package main

import "fmt"

func main() {
	age := 18
	if age <= 7 {
		fmt.Println("You're a child")
	} else if age <= 17 {
		fmt.Println("You're a teenager")
	} else {
		fmt.Println("You're an adult")
	}
}
