package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	fmt.Println("Write ", i, " as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("It's the afternoon")
	}
}
