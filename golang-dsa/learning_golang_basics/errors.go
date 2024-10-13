package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("no_name.txt")
	if err != nil{
		fmt.Println("Printing custom error : ", errors.New("I am a customer Error"), err)
		return
	}
	fmt.Println(file)
}
