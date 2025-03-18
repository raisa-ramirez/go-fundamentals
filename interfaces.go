package main

import (
	"fmt"
	"math"
)

type rect struct {
	base, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.base * r.height
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type geometry interface {
	area() float64
}

func result(g geometry) {
	fmt.Println(g)
	fmt.Println(math.Round(g.area()))
}

func main() {
	rect := rect{base: 4, height: 3}
	circle := circle{radius: 5}

	result(rect)
	result(circle)
}
