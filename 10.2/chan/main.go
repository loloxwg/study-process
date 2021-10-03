package main

import "time"

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <- c:
				println(v)
			case <- time.After(5 * time.Second):
				println("timeout")
				o <- true
				// 此处的break只是跳出了select循环，并未终止for循环，要用return才能终止这个子进程
				break
			}
		}
	}()
	<- o
}
