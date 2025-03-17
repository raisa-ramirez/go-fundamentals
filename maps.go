package main

import "fmt"

func main() {
	// First way
	// var cardData map[string]string = make(map[string]string)

	// Second way
	cardData := make(map[string]string)

	cardData["name"] = "Raisa"
	cardData["lastname"] = "Ramírez"
	cardData["country"] = "El Salvador"

	fmt.Println(cardData)
	fmt.Println(cardData["name"])

	fmt.Println(len(cardData))

	delete(cardData, "country")
	fmt.Println(cardData)

	clear(cardData)
	fmt.Println(cardData)
}
