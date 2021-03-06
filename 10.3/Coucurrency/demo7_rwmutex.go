package main

import (
	"fmt"
	"sync"
	"time"
)


var rwMutex *sync.RWMutex
var wgs *sync.WaitGroup
func main() {
	rwMutex = new(sync.RWMutex)
	wgs = new (sync.WaitGroup)

	//wgg.Add(2)
	//
	////多个同时读取
	//go readData(1)
	//go readData(2)

	wgs.Add(3)
	go writeData(1)
	go readData(2)
	go writeData(3)

	wgs.Wait()
	fmt.Println("main..over...")
}


func writeData(i int){
	defer wgs.Done()
	fmt.Println(i,"开始写：write start。。")
	rwMutex.Lock()//写操作上锁
	fmt.Println(i,"正在写：writing。。。。")
	time.Sleep(3*time.Second)
	rwMutex.Unlock()
	fmt.Println(i,"写结束：write over。。")
}

func readData(i int) {
	defer wgs.Done()

	fmt.Println(i, "开始读：read start。。")

	rwMutex.RLock() //读操作上锁
	fmt.Println(i,"正在读取数据：reading。。。")
	time.Sleep(3*time.Second)
	rwMutex.RUnlock() //读操作解锁
	fmt.Println(i,"读结束：read over。。。")
}
