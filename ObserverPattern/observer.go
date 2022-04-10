package main

type observer interface {
	NotifyObserver(temp float32)
}
