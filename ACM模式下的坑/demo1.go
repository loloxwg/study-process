package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		nums := strings.Split(input.Text(), " ") //分割字符串
		if nums[0] == "0" { // 判断是否结束
			break
		}

		hash:=map[float64]bool{}
		array:=[]float64{}
		for i := 0; i < len(nums); i++ {
			num, _ := strconv.ParseFloat(nums[i],32) // 字符串转数字

			if hash[num]==true {
				continue
			}
			hash[num]=true
			array=append(array,num)
		}
		fmt.Println(len(array))
		fmt.Println(array)
	}
}

