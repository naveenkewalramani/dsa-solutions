package main

import "fmt"

func main() {
	// create an array of desired size with each value initialize to 0
	var x [5]int
	fmt.Println(x)
	fmt.Println(len(x))
	// range over array
	for index, value := range x {
		fmt.Println(index, value)
	}

	// y:= type{values} //composite literal(allows values of same type to group together)
	y := []int{1, 2, 2, 3, 4}
	fmt.Println(y)
	fmt.Println(len(y))
	// range over slice
	for index, value := range y {
		fmt.Println(index, value)
	}

	z := []int{}
	fmt.Println(z)

	//slicing a slice
	a := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(a)
	// till end
	fmt.Println(a[0:])
	fmt.Println(a[1:])
	// 0 to n-1
	fmt.Println(a[:3])
	fmt.Println(a[0:4])

	// append slice
	b := []string{"a", "b", "c", "d"}
	fmt.Println(b)
	b = append(b, "n", "m", "o")
	fmt.Println(b)

	//delete from slice
	c := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(c)
	c = append(c[:2], c[4:]...)
	fmt.Println(c)

	// make([]type, lenght, capacity)
	d := make([]int, 10, 100)
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	//multidimensional string(slice of slice)
	f := [][]string{
		{"Naveen", "Kewalramani", "Male", "23"}, {"Ayushi", "Jaiswal", "Female", "22"}}
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))
	fmt.Println(f[1])
	fmt.Println(len(f[1]))
	fmt.Println(cap(f[1]))

}
