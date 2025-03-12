package channels

import (
	"fmt"
	"time"
)

//We define a PubSub struct that contains:
//subscribers → A map where each topic (string) has a list of subscriber channels ([]chan string).

type PubSub struct {
	subscribers map[string][]chan string
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
	ch := make(chan string, 5) // Buffered channel
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return ch
}

func (ps *PubSub) Publish(topic, message string) {
	for _, sub := range ps.subscribers[topic] {
		sub <- message
	}
}

func EventDrivenArchitecture() {
	ps := NewPubSub()

	sub1 := ps.Subscribe("news")
	sub2 := ps.Subscribe("news")

	go func() {
		for msg := range sub1 {
			fmt.Println("📢 Subscriber 1 received:", msg)
		}
	}()

	go func() {
		for msg := range sub2 {
			fmt.Println("📢 Subscriber 2 received:", msg)
		}
	}()

	ps.Publish("news", "Breaking: Go is awesome!")
	ps.Publish("news", "Update: Golang 2.0 released!")

	time.Sleep(time.Second)

}
