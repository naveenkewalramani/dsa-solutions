package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password1 := "uubrubruvbur"
	password2 := password1
	password3 := "randomString"
	// Password creation and compare with min cost
	bcryptPass1 := genPass(password1, bcrypt.MinCost)
	passCompare(password2, bcryptPass1)
	passCompare(password3, bcryptPass1)

	// Password creation and compare with default cost
	bcryptPass2 := genPass(password1, bcrypt.DefaultCost)
	passCompare(password2, bcryptPass2)
	passCompare(password3, bcryptPass2)

	// Cost method
	fmt.Println(bcrypt.Cost(bcryptPass1))
	fmt.Println(bcrypt.Cost(bcryptPass2))
}

// generate password hash with password and cost
func genPass(password string, cost int) []byte {
	// bcryptPass is the hash version for password
	bcryptPass, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err == nil {
		fmt.Println("Encrypt password hash : ", bcryptPass)
	} else {
		fmt.Println("Error in password Creation : ", err)
	}
	return bcryptPass
}

// compare password
func passCompare(inputPassword string, bcyrptPassword []byte) {
	response := bcrypt.CompareHashAndPassword(bcyrptPassword, []byte(inputPassword))
	if response == nil {
		fmt.Println("Password match")
	} else {
		fmt.Println("Password does not match : ", response)
	}
}
