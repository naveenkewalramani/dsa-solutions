package main

import "fmt"

type vehcile struct {
	doors int
	color string
}

type truck struct {
	vehcile
	fourWheel bool
}

type sedan struct {
	vehcile
	luxury bool
}

func main() {
	v1 := vehcile{
		doors: 4,
		color: "Black",
	}
	t1 := truck{
		vehcile:   v1,
		fourWheel: true,
	}
	s1 := sedan{
		vehcile: v1,
		luxury:  true,
	}
	fmt.Println(v1, v1.doors, v1.color)
	fmt.Println(t1, t1.vehcile, t1.vehcile.doors, t1.vehcile.color, t1.fourWheel)
	fmt.Println(s1, s1.vehcile, s1.vehcile.doors, s1.vehcile.color, s1.luxury)

}
