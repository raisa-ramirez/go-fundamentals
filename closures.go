package main

import "fmt"

func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	res := initSeq()
	fmt.Println(res())
	fmt.Println(res())
	fmt.Println(res())
	fmt.Println(res())
	fmt.Println(res())
}
