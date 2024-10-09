package orange

import "fmt"

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o Orange) String() string {
	return fmt.Sprintf("%v", o.Quantity)
}
