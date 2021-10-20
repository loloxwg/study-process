package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(time.Second)

		fmt.Println("Goroutine1 finished")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second)

		fmt.Println("Goroutine2 finished")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("All Goroutine finished")
}
