package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var tickett =10
var mutex sync.Mutex
var wgg sync.WaitGroup
func main() {
	wgg.Add(4)
	go saleTicketss("shop1")
	go saleTicketss("shop2")
	go saleTicketss("shop3")
	go saleTicketss("shop4")
	wgg.Wait()
	//time.Sleep(5*time.Second)
}
func saleTicketss(name string){

	rand.Seed(time.Now().UnixNano())

	defer wgg.Done()
	for{

		mutex.Lock()
		if tickett >0{
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name,"售出 ", tickett)

			tickett--

		}else {
			mutex.Unlock()
			fmt.Println(name,"卖完了")
			break
		}
		mutex.Unlock()
	}

}