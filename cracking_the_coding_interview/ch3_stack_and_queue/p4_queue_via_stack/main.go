package main

import "fmt"

func main() {
	q := queueConstructor()
	fmt.Println(q.Dequeue())
	fmt.Println(q.Peek())
	fmt.Println(q.IsEmpty())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Peek())
	fmt.Println(q.IsEmpty())

}
