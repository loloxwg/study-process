package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("writing data into file")
	writeFile("welcome to go")
	readFile()
	fmt.Println("read data into file")
	readFile()
}

func writeFile(message string) {
	bytes := []byte(message)
	ioutil.WriteFile("test.txt", bytes, 0644)
	fmt.Println("created a file")

}

func readFile() {
	data, _ := ioutil.ReadFile("test.txt")
	fmt.Println("file read from file")
	fmt.Println(string(data))
}
