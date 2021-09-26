package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统时候遍历时不关闭管道会造成阻塞 导致deadLock
	//问题是开发中不好确定什么时候关闭管道，
	//可以用select解决
	for {
		select {
		case v := <-intChan:
			fmt.Printf("intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("stringChan取得的数据是%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("全部取不到，程序员赶紧加逻辑\n")
			time.Sleep(time.Second)
			return
		}
	}

}
