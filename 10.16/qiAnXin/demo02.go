package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var char_list=[]int{32, 39, 45, 33, 34, 35, 36, 37, 38, 40, 41, 42, 44, 46, 47, 58, 59, 63, 64, 91, 92, 10, 9, 93, 94, 95, 96, 123, 124, 125, 126, 43, 60, 61, 62, 215, 247, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 65, 98, 66, 99, 67, 100, 68, 101, 69, 102, 70, 103, 71,104, 72, 105, 73, 106, 74, 107, 75, 108, 76, 109, 77, 110, 78, 111, 79, 112, 80, 113, 81, 114, 82, 115, 83, 116, 84, 117, 85, 118, 86,119, 87, 120, 88, 121, 89, 122, 90}

func main() {
	//a:="1 2 3 4 5"
	//b:="1"
	////c:="1.0"
	//n1,_:=strconv.Atoi(b)
	//fmt.Println(n1)
	//n2,_:=strconv.ParseFloat(b,32)///
	//fmt.Println(n2)
	input := bufio.NewScanner(os.Stdin) //创建并返回一个从os.Stdin读取数据的Scanner
	for input.Scan(){
		// Scan方法获取当前位置的token（该token可以通过Bytes或Text方法获得），
		// 并让Scanner的扫描位置移动到下一个token。
		// 当扫描因为抵达输入流结尾或者遇到错误而停止时，本方法会返回false
		nums := strings.Split(input.Text(), ",") //分割字符串
		if nums[0] == "0" { // 判断是否结束
			break
		}
		array:=[]int{}
		for i := 0; i < len(nums); i++ {
			num, _ := strconv.ParseFloat(nums[i],32) // 字符串转数字

			c:=swapMap(int(num))
			array=append(array,c)
		}
		for i := 0; i < len(nums); i++ {
			if i==len(nums)-1 {
				fmt.Printf("%d\n",array[i])
			}else {
				fmt.Printf("%d,",array[i])
			}

		}

	}
}
func swapMap(a int) int {
	hmap:=map[int]int{}
	for i := 0; i < 99; i++ {
		hmap[char_list[i]]=i+1
	}
	c:=hmap[a]
	return c
}