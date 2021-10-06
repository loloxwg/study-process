package main

import "fmt"

//混合工厂模式
type Animal struct {
	species string
	age     int
}

type AnimalHouse struct {
	name         string
	sizeInMeters int
}

type AnimalFactory struct {
	species   string
	houseName string
}

// NewAnimal is an `Animal` factory that fixes
// the species as the value of its `AnimalFactory` instance
func (af AnimalFactory) NewAnimal(age int) Animal{
	return Animal{
		species: af.species,
		age:     age,
	}
}

// NewHouse is an `AnimalHouse` factory that fixes
// the name as the value of its `AnimalFactory` instance
func (af AnimalFactory) NewHouse(sizeInMeters int) AnimalHouse{
	return AnimalHouse{
		name:         af.houseName,
		sizeInMeters: sizeInMeters,
	}
}


func  main() {
	dogFactory := AnimalFactory{
		species: "dog",
		houseName: "kennel",
	}
	dog:=dogFactory.NewHouse(2)
	kennel:=dogFactory.NewAnimal(3)
	fmt.Println("dog",dog,"kennel",kennel)

	horseFactory := AnimalFactory{
		species: "horse",
		houseName: "stable",
	}
	horse := horseFactory.NewAnimal(12)
	stable := horseFactory.NewHouse(30)
	fmt.Println("horse",horse,"stable",stable)
}