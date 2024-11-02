package main

import "fmt"

func main() {
	favSport := "Football"
	switch favSport {
	case "Cricket":
		fmt.Println("it is cricket")
	case "Football":
		fmt.Println("it is football")
	case "BasketBall":
		fmt.Println("it is basketball")
	}
}
