package main

import (
	"fmt"
)

type CustomErr struct {
	info string
}

func (ce CustomErr) Error() string {
	return fmt.Sprintf("here is the error : %v", ce.info)
}


func main(){
	customError := CustomErr{info: "File Not Present Exception"}
	foBar(customError)
}

func foBar (e error) {
	fmt.Println(e)
	fmt.Println(e.(CustomErr).info) // (CustomErr). this thing is called assertion
}
