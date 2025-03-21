package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	figure := rect{width: 10, height: 5}
	fmt.Println("Area: ", figure.area())
	fmt.Println("Perimetro: ", figure.perim())
}
