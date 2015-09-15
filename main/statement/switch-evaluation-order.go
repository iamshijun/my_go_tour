package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand_3(seed int64) int64 {
	rand.Seed(seed) // seed must be int64
	return rand.Int63n(3)
}

func main() {
	fmt.Println("When's Sunday?")
	today := time.Now().Weekday()
	//fmt.Printf("time.Sunday %T\n", time.Sunday) //Type is "time.Weekday"
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	switch i := rand.Int63n(3); i {
	case rand_3(i):
		fmt.Println("Haha , You guess the right num")
	default:
		fmt.Println("Sorry, You miss the num", i)
	}

}
