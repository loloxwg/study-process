package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := "test.txt"
	destFile := "bbb2.txt"
	total,err:=CopyFile2(srcFile,destFile)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(total)

}
func CopyFile(src, dest string) (int, error) {
	file1, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer func(file2 *os.File) {
		err := file2.Close()
		if err != nil {

		}
	}(file2)
	defer func(file1 *os.File) {
		err := file1.Close()
		if err != nil {

		}
	}(file1)

	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0
	defer func() {
		recover()
	}()
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("end of file copy")
			return total,nil
		}else if err != nil {
			fmt.Println("error reading file")
			panic("panic")
		}
		total+=n
		file2.Write(bs[:n])
	}
	return total,nil
}

func CopyFile2(src,dest string)(int64 , error){
	file1,err:=os.Open(src)
	if err != nil {
		return 0,err
	}
	file2,err:=os.OpenFile(dest,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil {
		return 0,err
	}
	defer func() {
		err := file1.Close()
		if err != nil {
			return
		}
	}()
	defer func() {
		err := file2.Close()
		if err != nil {
			return
		}
	}()

	return io.Copy(file2,file1)

}

