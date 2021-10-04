package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket=10
var mu sync.Mutex
func main() {
	go saleTickets("shop1")
	go saleTickets("shop2")
	go saleTickets("shop3")
	go saleTickets("shop4")

	time.Sleep(5*time.Second)
}
func saleTickets(name string){
	mu.Lock()
	rand.Seed(time.Now().UnixNano())
	for{
		if ticket >0{
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name,"售出 ",ticket)

			ticket--

		}else {
			fmt.Println(name,"卖完了")
			break
		}
	}
	mu.Unlock()
}