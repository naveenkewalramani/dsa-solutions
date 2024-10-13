package main

import "fmt"

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func info(s shape) {
	a := s.area()
	fmt.Println(a)
}

func main() {
	sq := square{side: 10}
	ci := circle{radius: 2}
	fmt.Println(sq.area(), ci.area())
	info(sq)
	info(ci)
}
