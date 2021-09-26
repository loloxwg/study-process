package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]uint64, 10)
	//声明全局互斥锁
	//lock 是一个全局互斥锁
	//sync 是包：synchronized
	//mutex 是互斥
	lock sync.Mutex
)

func test1(n int) {
	sum := uint64(1)
	for i := 1; i <= n; i++ {
		sum += uint64(i)
	}
	lock.Lock()
	//加锁
	myMap[n] = sum
	//解锁
	lock.Unlock()
}

//func test2(n int) int {
//	_, b, r := 0, 1, 2
//	for i := 0; i < n; i++ {
//		_, b, r = b, r, b*r
//	}
//	return r
//}

func main() {
	for i := 1; i <= 200; i++ {
		go test1(i)
	}

	time.Sleep(time.Second * 10)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d] = %v\n", i, v)
	}
	lock.Unlock()

}
