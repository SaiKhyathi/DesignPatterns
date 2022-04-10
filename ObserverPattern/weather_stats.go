package main

var temp float32
var humidity float32
var pressure float32
var observers []observer

func registerObserver(observer observer) {
	observers = append(observers, observer)
}

func removeObserver(observer observer) {
	for i := 0; i < len(observers); i++ {
		if observer == observers[i] {
			observers = append(observers[:i], observers[i+1:]...)
			break
		}
	}
}

func notifyObservers(temp float32) {
	for i := 0; i < len(observers); i++ {
		observers[i].NotifyObserver(temp)
	}
}

func main() {
	conditions := CurrentConditions{}
	registerObserver(conditions)
	temp = 123.3
	notifyObservers(temp)
	removeObserver(conditions)
}
