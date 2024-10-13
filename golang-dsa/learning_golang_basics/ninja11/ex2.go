package main
import (
	"encoding/json"
	"fmt"
)
type User struct {
	name            string
	email           string
	phone           string
	accountNumber   string
	amount string
	currency string
}


func main() {
	user1 := User{"Naveen Kewalramani",
		"naveen.kewalramani2903@gmail.com",
		"81", "222", "1991", "INR" }
	fmt.Println(user1)
	m, err := json.Marshal(user1)
	fmt.Errorf("this is an error 1")
	if err != nil{
		fmt.Errorf("this is an error 2")
		return
	}
	fmt.Println(m)
}