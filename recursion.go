package main

import "fmt"

func fac(num int) int {
	if num == 0 {
		return 1
	}
	return num * fac(num-1)
}

func main() {
	res := fac(4)
	fmt.Println(res)
}
