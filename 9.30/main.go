package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.baidu.com",
		"http://www.jd.com",
		"http://www.taobao.com",
		"http://www.163.com",
		"http://www.sohu.com",
	}

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}
	fmt.Println(link, "is up")
}
