package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println("Using switch..true statement.")

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	fmt.Println("Using if..elseif..else statement.")

	if t.Hour() < 12 {
		fmt.Println("Good morning!")
	} else if t.Hour() < 17 {
		fmt.Println("Good afternoon.")
	} else {
		fmt.Println("Good evening.")
	}
}

/*
	没有条件的 switch 同 switch true 一样。
	这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
*/
