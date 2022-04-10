package main

var cf ChocolateFactory = ChocolateFactory{
	Empty:  false,
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
		cf.Boiled = false
	}
}

func boil() {
	if !isEmpty() && !isBoiled() {
		cf.Boiled = true
	}
}

func main() {

}
