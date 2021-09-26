package main

import (
	"fmt"
)

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool, count int) {
	var flag bool
	for {
		//time.Sleep(time.Millisecond*10)
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true //假设是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}

	}

	fmt.Printf("第%d个primeNumber 协程因为取不到数据而退出\n", count)
	exitChan <- true

}
func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	go putNum(intChan)

	for i := 0; i < 4; i++ { //开四个协程去判断
		go primeNum(intChan, primeChan, exitChan, i+1)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数=", res)
	}
	fmt.Println("主线程退出")

}
