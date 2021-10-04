package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main() {
	/*同步等待组*/

	wg.Add(2)//+2
	go fun1()
	go fun2()

	fmt.Println("main()进入阻塞状态。。。等待wg中的子goroutine结束 blocked...")
	wg.Wait()
	fmt.Println("main goroutine unlocked")
}

func fun1()  {
	for i := 0; i <10; i++ {
		fmt.Println("fun1()中打印...A",i)
	}
	wg.Done()//-1
}

func fun2()  {
	defer wg.Done()//-1
	for i := 0; i <10; i++ {
		fmt.Println("\tfun2()中打印...B",i)
	}

}
