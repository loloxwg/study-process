package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Printf(
		"&r1.leftUp.x=%p\n&r1.leftUp.y=%p\n"+
			"&r1.rightDown.x=%p\n&r1.rightDown.y=%p\n",
		&r1.leftUp.x,
		&r1.leftUp.y,
		&r1.rightDown.x,
		&r1.rightDown.y,
	)
	r2 := Rect2{&Point{1, 2}, &Point{3, 4}}
	fmt.Printf(
		"r2.leftUp=%p\nr2.rightDown=%p\n指向的地址\n",
		r2.leftUp,
		r2.rightDown,
	)
	fmt.Printf(
		"r2.leftUp=%p\nr2.rightDown=%p\n本身的地址\n",
		&r2.leftUp,
		&r2.rightDown,
	)
}
