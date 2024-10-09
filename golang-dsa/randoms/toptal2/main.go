package main

import (
	"fmt"
	"toptal2/queue"
)

func main() {
	q := queue.IQueue{}
	q.Init()
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Dequeue())
	q.Enqueue(10)
	q.Enqueue(20)
	fmt.Println(q.Dequeue())
	fmt.Println(q.IsEmpty())
}
