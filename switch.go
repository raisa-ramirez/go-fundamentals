package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Weekday()

	switch day {
	case time.Monday:
		fmt.Print("You're starting the week...")
	case time.Tuesday, time.Wednesday, time.Thursday:
		fmt.Print("You're working...")
	case time.Friday:
		fmt.Print("You're finishing your work...")
	default:
		fmt.Print("You're resting...")
	}
}
