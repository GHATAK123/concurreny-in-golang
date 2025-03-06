package main

import (
	"concurrency-in-go/channels"
	"concurrency-in-go/closure"
	"concurrency-in-go/generics"
	"concurrency-in-go/ood"
	waitgroups "concurrency-in-go/waitGroups"
	"fmt"
)

func main() {
	fmt.Println("Concurrency In Golang!!!")
	waitgroups.Helper()
	channels.Helper()
	closure.Helper()
	generics.Helper()
	ood.Helper()
	ood.CompositionHelper()
	ood.PolyHelper()
	ood.EncapsulationHelper()
}
