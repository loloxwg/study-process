package main

import (
	"fmt"
	"io"
)

///1 2   1*1+1
///2 5   2*2+1
///3 10
///4 17

func main() {
	a:=0
	for {
		_, err := fmt.Scanf("%v\n",&a)
		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			m,res:=hhh(a)
			fmt.Println(m,res)
		}
	}
}


func hhh(n int)(m ,res int) {
	array:=make([]int,n)
	var sum int

	for i := 0; i < len(array); i++ {
		if sum>n{
			break
		}
		array[i]=(i+1)*(i+1)+1
		sum+=array[i]
		m=i


	}
	res=n-array[m+3]

	return m,res-17

}