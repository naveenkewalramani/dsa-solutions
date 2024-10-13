package main

import "fmt"

func main() {
	// slice
	a := []int{1, 2, 3}
	fmt.Println(a)

	//map(key type string)(value type int)
	b := map[string]string{
		"James Bond":        "21",
		"Navee Kewalramani": "32",
	}
	b["Ayushi"] = "11"
	fmt.Println(b)
	fmt.Println(b["Ayushi"])
	fmt.Println(b["James Bond"])
	fmt.Println(b["JamesBond"])
	_, ok := b["James Bond"]
	if ok {
		fmt.Println("Value Exist")
	} else {
		fmt.Println("Not Exist")
	}

	c := map[string]string{
		"naveen": "292929",
		"tina":   "2929232",
	}
	fmt.Println(c["22"])

	for index, value := range c {
		fmt.Println(index, value)
	}
	delete(c, "naveen")
	delete(c, "naveen1")
	fmt.Println(c)
}
