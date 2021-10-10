package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
//滴滴笔试
func main()  {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	t, _ := strconv.Atoi(input.Text())
	fmt.Println(t)//

	input2:=bufio.NewScanner(os.Stdin)
	input2.Scan()
	nums:=strings.Split(input2.Text()," ")
	fmt.Println(nums)


	hash:=map[float64]bool{}
	array:=[]float64{}
	for i := 0; i < t; i++ {
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
