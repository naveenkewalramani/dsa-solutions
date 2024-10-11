package generics

import "golang.org/x/exp/constraints"

func Addition[Num constraints.Ordered](a ...Num) Num {
	var sum Num
	for _, v := range a {
		sum += v
	}
	return sum
}

func Subtraction[Num constraints.Integer](a, b Num) Num {
	return a - b
}
