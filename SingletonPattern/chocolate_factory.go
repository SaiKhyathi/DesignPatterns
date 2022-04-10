package main

import "fmt"

//Init the chocolate factory at the startup so that it is not re-initiaed again
var cf ChocolateFactory = ChocolateFactory{
	Empty:  true,
	Boiled: false,
}

type ChocolateFactory struct {
	Empty  bool
	Boiled bool
}

func isEmpty() bool {
	return cf.Empty
}

func isBoiled() bool {
	return cf.Boiled
}

func fill() {
	if cf.Empty {
		cf.Empty = false
		fmt.Println("Vessel is filling")
	} else {
		fmt.Println("Vessel already full")
	}
}

func boil() {
	if !isEmpty() && !isBoiled() {
		cf.Boiled = true
		fmt.Println("Boiling")
	} else if isEmpty() {
		fmt.Println("Can't boil empty vessel")
	} else {
		fmt.Println("Already Boiled")
	}
}

func emptyVessel() {
	if !isEmpty() {
		cf.Empty = true
		cf.Boiled = false
		fmt.Println("Emptying the vessel")
	} else {
		fmt.Println("Already Empty")
	}
}

func main() {
	fmt.Println("Boiling ", cf.Boiled)
	fmt.Println("Empty Vessel ", cf.Empty)
	boil()
	fill()
	fmt.Println("Post filling state of Boiling ", cf.Boiled)
	fmt.Println("Post filling state of Empty Vessel", cf.Empty)
	fill()
	boil()
	fmt.Println("Post boiling state of Boiling ", cf.Boiled)
	fmt.Println("Post boiling state of Empty Vessel ", cf.Empty)
	boil()
	emptyVessel()
	boil()
	fill()
	boil()
}
