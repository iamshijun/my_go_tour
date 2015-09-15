package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	//see we can get rid of the
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Windows.")
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		//freebsd , openbds,plan9,windows...
		fmt.Printf("%s..", os)
	}
}
