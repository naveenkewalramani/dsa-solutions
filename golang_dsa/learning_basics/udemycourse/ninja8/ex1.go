package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	p1 := user{
		Name:   "Person 1",
		Age:    22,
		Gender: "Female",
	}

	p2 := user{
		Name:   "Person 2",
		Age:    24,
		Gender: "Male",
	}

	// Normal Hash Map
	p3 := map[string]int{
		"naveen": 1,
		"test2":  2,
	}

	fmt.Println(p1, p2, p3)

	// JSON Encoding
	j1, err := json.Marshal(p1)
	fmt.Println(j1, err)
	j2, err := json.Marshal(p2)
	fmt.Println(j2, err)
	j3, err := json.Marshal(p3)
	fmt.Println(j3, err)

	var p4 user
	var p5 map[string]int
	json.Unmarshal(j1, &p4)
	fmt.Println(p4)
	json.Unmarshal(j2, &p4)
	fmt.Println(p4)
	json.Unmarshal(j3, &p5)
	fmt.Println(p5)

}
