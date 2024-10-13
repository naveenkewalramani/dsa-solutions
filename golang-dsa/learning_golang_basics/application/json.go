package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name    string
	Age     int
	Hobbies []string
}

func main() {
	p1 := person{
		Name:    "Naveen",
		Age:     20,
		Hobbies: []string{"Cricket", "Table Tennis", "COD"},
	}
	p2 := person{
		Name:    "Niyati",
		Age:     27,
		Hobbies: []string{"Table Tennis"},
	}
	p3 := []person{p1, p2}

	// json Marshal
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	jsonData1, _ := json.Marshal(p1)
	jsonData2, _ := json.Marshal(p2)
	jsonData3, _ := json.Marshal(p3)
	fmt.Println(string(jsonData1))
	fmt.Println(string(jsonData2))
	fmt.Println(string(jsonData3))

	// json Unmarshal
	un1 := person{}
	un2 := person{}
	un3 := []person{}
	json.Unmarshal(jsonData1, &un1)
	json.Unmarshal(jsonData2, &un2)
	json.Unmarshal(jsonData3, &un3)
	fmt.Println(un1)
	fmt.Println(un2)
	fmt.Println(un3)

}
