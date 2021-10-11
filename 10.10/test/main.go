package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input1, _ := inputReader.ReadString('\n')
	n, _ := strconv.Atoi(input1)
	fmt.Printf(input1) //
	fmt.Println(n)

	inputReader2 := bufio.NewReader(os.Stdin)
	input2, _ := inputReader2.ReadString('\n')
//	fmt.Println(input2) //

	var context = strings.Fields(input2)
	for i := 0; i <n; i++ {
		digit, _ := strconv.Atoi(context[i])
		fmt.Printf("%d",digit)
	}
	//fmt.Println(context[1])
	//digit:=strconv.Atoi()
	//var digits =make([]int,n)
	//for i := 0; i < n; i++ {
	//	digit, _ := strconv.Atoi(context[i])
	//	fmt.Println(digit)
	//	b := digit
	//	for _,v :=range digits{
	//		if b!=v{
	//			digits = append(digits, b)
	//		}
	//	}
	//}
	//fmt.Println(len(digits))

}
