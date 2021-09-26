package main

import "fmt"

func main() {
	intChan := make(chan int, 3)

	intChan <- 100000
	intChan <- 123456789
	close(intChan)

	n1 := <-intChan

	fmt.Printf("n1:=%v\n", n1)

	intChan2 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan2 <- i
	}

	close(intChan2) //写入关闭再进行遍历

	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
