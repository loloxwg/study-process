package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {



	fileName:="test.txt"
	file,err:=os.OpenFile(fileName,os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bs:=[]byte{0}
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4,io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(2,0)
	read, err := file.Read(bs)
	if err != nil {
		return
	}
	fmt.Println(read)
}
