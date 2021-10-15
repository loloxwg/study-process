package main

import (
	"fmt"
	"io"
)

func main() {
	a:=""
	b:=""
	for {
		_, err := fmt.Scanf("%s\n%s\n",&a,&b)
		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			c:=reoder(a,b)
			fmt.Printf("%s\n",c )
		}
	}
}

func reoder(a, b string)string{
	if b == "" {
		return a
	}
	if a=="" {
		return ""
	}
	lenb:=len(b)
	m:=a[lenb:]
	n:=[]byte(b)
	nums:=append(n,m...)
	nu:=string(nums)
	return nu
}
