package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(x)
	for k, v := range x {
		fmt.Println("The key is : ", k, "and the value stored is: ", v)
	}
}
