package main

import (
	"fmt"
	"unsafe"
)

func main() {
	buitInTypeQuestions()
	compositeTypeQuestions()

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

	testFunc := func(a, b int) int {
		return a + b
	}
	fmt.Println(testFunc(1, 1))
	fmt.Println("Size of function", unsafe.Sizeof(testFunc))
}
