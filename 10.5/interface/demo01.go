package main

import "fmt"

type Animal interface {
	makeNoise() string
}

type Dog struct {
	name string
	legs int
}

func (d *Dog) makeNoise() string {
	return d.name + " says woof!"
}

type Duck struct {
	name string
	legs int
}

func (d *Duck) makeNoise() string {
	return d.name + " says quack!"
}

func NewDog(name string) Animal {
	return &Dog{
		legs: 4,
		name: name,
	}
}

func NewDuck(name string) Animal {
	return &Duck{
		legs: 4,
		name: name,
	}
}

func main() {
	var dog, duck Animal

	dog = NewDog("fido")
	duck = NewDuck("donald")

	fmt.Println(dog.makeNoise())
	// fido says woof!
	fmt.Println(duck.makeNoise())
	// donald says quack!
}
