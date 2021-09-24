package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr3 [5]int
	len := len(intArr3)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(1000)
	}
	fmt.Println(intArr3)

	for i := 0; i < len/2; i++ {
		//		intArr3[len-1-i],intArr3[len-1-i-1]=intArr3[i],intArr3[i+1] //TODO 交换元素
		tem := intArr3[len-i-1]
		intArr3[len-i-1] = intArr3[i]
		intArr3[i] = tem

	}
	fmt.Println(intArr3)
}
