package main

import (
	"fmt"
	"unsafe"
)

func main() {
	buitInTypeQuestions()
	compositeTypeQuestions()
	functionRelatedInterviewQuestions()
	fmt.Println(test([]string{"flower", "flow", "flowing", "follower"}))
}

func buitInTypeQuestions() {
	// built-in types
	// Numeric - int, int64, int32, int16, int8, unint, unint64, unint32, unint16, unint8, float64, float32, complex64, complex128
	// String - string
	// Boolean
	// rune - int32, byte - unint8 // alias of builtIn types

	typeOf("fuffufi")
	typeOf(32)
	typeOf(32.44)
}

func typeOf(variable interface{}) {
	switch variable.(type) {
	case int:
		fmt.Println("this is int")
	case float32, float64:
		fmt.Println("this is float")
	case string:
		fmt.Println("this is string")
	case bool:
		fmt.Println("this is string")
	default:
		fmt.Println("Invalid variable")
	}
}

func compositeTypeQuestions() {
	// pointer
	// functions
	// channels
	// interfaces
	// container - // slices, maps, array
	var ptr *int
	a := 30
	ptr = &a
	fmt.Println("Pointer to int, size", unsafe.Sizeof(ptr))
	s := struct {
		Age  int
		Flag bool
	}{
		20, true,
	}

	// custom struct
	fmt.Println("Custom struct size", unsafe.Sizeof(s))
	unbufferedCh := make(chan int)
	fmt.Println("Size of unbuffered channel", unsafe.Sizeof(unbufferedCh))
	bufferedCh := make(chan int, 19)
	fmt.Println("Size of buffered channel", unsafe.Sizeof(bufferedCh))

	testFunc := func(a, b int) int { // anonymous function or function literal
		return a + b
	}
	fmt.Println(testFunc(1, 1))
	fmt.Println("Size of function", unsafe.Sizeof(testFunc))
	// assiging function to another variable
	m := testFunc
	fmt.Println(m(1, 1))
}

func functionRelatedInterviewQuestions() {
	fmt.Println(computeAnotherFunction(add, 10, 5)) // passing function as variable and passing function to another function
	fmt.Println(computeAnotherFunction(sub, 10, 5))
	fmt.Println(computeAnotherFunction(mul, 10, 5))
	fmt.Println(computeAnotherFunction(div, 10, 5))
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}
func computeAnotherFunction(a func(int, int) int, b, c int) int {
	return a(b, c) // assigning function to variable
}

// ordering of mesurement of time
// nanosecond -> microsecond -> millisecond -> decisecond -> centisecond -> second
