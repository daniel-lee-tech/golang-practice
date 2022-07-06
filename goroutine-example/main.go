package main

import (
	"fmt"
	"time"
)

// a very simple function that we'll
// make async later on
func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Goroutine Tutorial")

	go compute(10)
	go compute(10)

	var input string
	fmt.Scanln(&input)
}
