package main

import (
	"fmt"
	"time"
)

func print(message string) {
	for i := 0; i < 3; i++ {
		fmt.Println(message, " ", i)
	}
}

func main() {
	print("Hola")

	go print("Raisa")
	time.Sleep(time.Second)
}
