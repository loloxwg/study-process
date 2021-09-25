package main

import (
	"fmt"
	"io"
)

//
//import "fmt"
//
//var num=[13]int{0,3,7,11,15,20}
//var r=[13]int{0}
//func maxMoney(n int) int{
//	if r[n]!=0 {
//		return r[n]
//	}
//	if n==0{
//		return 0
//	}
//	q:=num[n]
//	for i := 0; i < n; i++ {
//		q=max(q,r[i]+maxMoney(n-i))
//		r[n]=q
//	}
//	return q
//}
//
//func max(a,b int)int{
//	if 	a<b{
//		return b
//	}
//	return a
//}
//func main()  {
//	current :=12
//	money := maxMoney(current)
//	fmt.Println(money)
//}

func main() {
	a := 0
	b := 0
	for {
		_, err := fmt.Scanf("%d%d", &a, &b)
		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			fmt.Printf("%d\n", a+b)
			/////

		}
	}
}
