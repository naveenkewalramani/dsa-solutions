package main

type Stack interface {
	Push(element int)
	Pop() int
	Top() int
	IsEmpty() bool
}
