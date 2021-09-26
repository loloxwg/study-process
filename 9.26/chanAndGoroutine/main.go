package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)

//	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("writeData: ", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData: ", v)
	}
	exitChan <- true
	close(exitChan)
}
