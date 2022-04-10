package main

import "fmt"

type singletonInstance struct {
}

var instance *singletonInstance

func getInstance() *singletonInstance {

	if instance == nil {
		fmt.Println("Created new instance")
		instance = new(singletonInstance)
	} else {
		fmt.Println("Existing instance")
	}
	return instance
}

func main() {

	for i := 0; i < 10; i++ {
		getInstance()
	}

}
