package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {

	start := time.Now()
	num := runtime.NumCPU()
	fmt.Println(num)
	runtime.GOMAXPROCS(num)
	go test()
	go test()
	go test()
	for i := 0; i < 10; i++ {

		fmt.Println("main() hello golang" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
	end := time.Now()
	time_sum := end.Sub(start)

	fmt.Println(time_sum)

}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() hello world" + strconv.Itoa(i))
		//time.Sleep(1*time.Second)
	}
}

//并发确实不安全，如果主协程执行完毕了，子协程还未完成，直接强制结束
