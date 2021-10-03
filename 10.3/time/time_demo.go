package main

import (
	"fmt"
	"time"
)

func main() {
	//时间戳
	now:=time.Now()
	timeStamp := now.Unix()
	fmt.Println(timeStamp)
}
