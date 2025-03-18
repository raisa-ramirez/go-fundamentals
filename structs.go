package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	userOne := person{name: "Raisa", age: 30}
	userTwo := person{name: "Nataren"}
	fmt.Println(userOne.name)
	fmt.Println(userOne.age)

	fmt.Println(userTwo.name)
	fmt.Println(userTwo.age)

}
