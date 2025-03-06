package main

import (
	"concurrency-in-go/channels"
	waitgroups "concurrency-in-go/waitGroups"
	"fmt"
)

func main() {
	fmt.Println("Concurrency In Golang!!!")
	waitgroups.Helper()
	channels.Helper()
}
