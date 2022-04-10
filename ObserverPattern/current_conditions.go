package main

import "fmt"

type CurrentConditions struct {
	tempCurrentConditions float32
}

func (c CurrentConditions) NotifyObserver(temp float32) {
	fmt.Println(temp, " :: temperature@CurrentConditions")
	c.tempCurrentConditions = temp
}
