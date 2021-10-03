package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	fileName:="D:\\xwg\\Documents\\资料\\-DailyStudy\\test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	b1 := bufio.NewReader(file)
	p:=make([]byte,1024)
	n1, err := b1.Read(p)
	fmt.Println(n1)
	fmt.Println(string(p[:n1]))




}
