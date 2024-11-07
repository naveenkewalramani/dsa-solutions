package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	author      string
	description string
	createdAt   string
	done        bool
}

func main(){
	todo1 := Todo {
		author: "Naveen Kewalramani",
		description: "Need to cut Nails",
		createdAt: "2021-04-21",
		done: false,
	}
	encodeTodo, err := json.Marshal(todo1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(encodeTodo)
	fmt.Println(todo1)
}
