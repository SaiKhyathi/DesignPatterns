package main

import "fmt"

type Forecast struct {
	tempForecast float32
}

func (f Forecast) NotifyObserver(temp float32) {
	fmt.Println(temp, " :: temperature@Forecast")
	f.tempForecast = temp
}
