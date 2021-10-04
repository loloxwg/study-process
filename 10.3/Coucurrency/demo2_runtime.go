package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("GOROOT",runtime.GOROOT())
	fmt.Println(
		runtime.GOOS,
		runtime.NumCPU(),
		runtime.GOMAXPROCS(8),

	)

	//go func() {
	//	for i:=0;i<5; i++{
	//		fmt.Println("goroutine...")
	//	}
	//}()
	//
	//for i := 0; i <4; i++ {
	//	runtime.Gosched()
	//	fmt.Println("main...")
	//}
	go func() {
		fmt.Println("goroutine...start...")
		fun()
		fmt.Println("goroutine...end...")
	}()
	time.Sleep(time.Second*1)
}

func fun(){
	defer fmt.Println("defer...")
	//return //终止函数
	runtime.Goexit()//终止当前goroutine
	fmt.Println("fun()...")
}