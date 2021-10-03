package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(
		fileInfo.Name(),
		fileInfo.Size(),
		fileInfo.Mode(),
		fileInfo.IsDir(),
		fileInfo.Sys(),
	)
	fileName1:="aa.txt"
	file3, err := os.Open(fileName1)
	if err != nil {
		errors.New("xxxxx ")
	}
	file3.Readdirnames(36)

	file3.Close()
	err=os.RemoveAll("/aaa")
	if err != nil {
		fmt.Println(err)
		errors.New("ssss")
	}




}