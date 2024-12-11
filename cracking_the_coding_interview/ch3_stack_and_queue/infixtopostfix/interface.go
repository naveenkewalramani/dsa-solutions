package main

type Stack interface {
	Push(element string)
	Pop() string
	Top() string
	IsEmpty() bool
}
