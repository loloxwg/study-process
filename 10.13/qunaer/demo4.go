package main

import (
	"fmt"
	"io"
)

func main() {
	a:=[]int{}
	b:=0
	c:=[]int{}
	for {
		_, err := fmt.Scanf("%v\n%d\n%v\n",&a,&b,&c)
		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			d:=pian(a,b,c)
			fmt.Printf("%d\n",d )
		}
	}
}


func pian(a []int ,b int ,c []int)int{
	return 1
}
