package main

import "fmt"

type Cat struct {
	name string
	age  int
}

func main() {

	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan 的值=%v "+
		"本身的地址=%p\n", intChan, &intChan)
	intChan <- 1000
	num := 211
	intChan <- num
	intChan <- 500

	fmt.Printf("channel len: %d\n"+
		"channel cap: %d\n",
		len(intChan), cap(intChan))

	var num2 int
	num2 = <-intChan
	fmt.Println("num2: ", num2)
	fmt.Printf("channel len: %d\n"+
		"channel cap: %d\n",
		len(intChan), cap(intChan))

	num3 := <-intChan
	num4 := <-intChan
	//num5:=<-intChan//deadLock
	//fmt.Println("num3: ",num3,"num4: ",num4,"num5: ",num5)
	fmt.Println("num3: ", num3, "num4: ", num4)

	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 10)
	m1["city1"] = "beijing"
	m1["city2"] = "shanghai"
	m2 := make(map[string]string, 20)
	m2["hero1"] = "songjiang"
	m2["hero2"] = "wusong"

	mapChan <- m1
	mapChan <- m2

	mapp1 := <-mapChan
	mapp2 := <-mapChan
	fmt.Printf("mapp1=%v\n"+
		"mapp2=%v\n", mapp1, mapp2)

	var catChan chan Cat

	catChan = make(chan Cat, 10)

	cat1 := Cat{name: "tom", age: 1}
	cat2 := Cat{name: "som", age: 2}

	catChan <- cat1 //放进去
	catChan <- cat2

	cat11 := <-catChan
	cat22 := <-catChan

	fmt.Printf("cat1=%v\ncat2=%v\n", cat11, cat22)

	var catChan2 chan *Cat
	catChan2 = make(chan *Cat, 10)

	catChan2 <- &cat1
	catChan2 <- &cat2

	cat111 := <-catChan2
	cat222 := <-catChan2
	fmt.Println(cat111, cat222)

	//任意数据类型
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	allChan <- cat1
	allChan <- cat2
	allChan <- 1000
	allChan <- "jack"
	allChan <- m1

	cat1111 := <-allChan
	fmt.Println(cat111.age)
	fmt.Printf("cat1111=%T,cat111=%v\n", cat1111, cat1111)

	a := cat1111.(Cat)
	fmt.Printf("cat111.name=%v\n", a.name)

	cat2222 := <-allChan
	v1 := <-allChan
	v2 := <-allChan
	v3 := <-allChan

	fmt.Println(cat1111, cat2222, v1, v2, v3)

}
