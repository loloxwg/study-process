package main

import "fmt"

func main() {

	////只写
	//var chan2 chan <-int
	//chan2=make(chan int ,3)
	//chan2<-200
	////num:=<-chan2 //error
	//fmt.Println("chan2: ",chan2)
	//
	////只读
	//var chan3 <-chan int
	//chan3=make(chan int ,3)
	//num2:=<-chan3
	////chan3<-11111111111111111111//error
	//fmt.Println("num2: ",num2)

	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("ending")

}

//只读
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)

	}
	var a struct{}
	exitChan <- a

}

//只写
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i

	}
	close(ch)
	var a struct{}
	exitChan <- a
}
