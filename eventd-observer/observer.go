package main

type Observer interface {
	Update(event string, data interface{})
}

type Subject interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	notify(event string, data interface{})
}
