package main

import "fmt"

func vals() (int, string) {
	return 30, "Raisa"
}

func main() {
	age, name := vals()
	fmt.Println("Age ", age)
	fmt.Println("Name ", name)
}
