package main

import (
	"fmt"
	"time"
)

func main() {
	go printNumber()

	for i := 0; i < 1000; i++ {
		fmt.Printf("\t main goroutine A\n")
	}
	time.Sleep(100 * time.Millisecond)
}

func printNumber() {
	for i := 0; i < 1000; i++ {
		fmt.Printf("child goroutine %d\n",i)

	}
}
