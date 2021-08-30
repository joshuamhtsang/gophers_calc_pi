package main

import (
	"fmt"

	"github.com/joshuamhtsang/gophers_calc_pi/geometry"
)

func main() {
	fmt.Println("Gophers Calculate Pi!")

	fmt.Println(geometry.WithinUnitCircle(0.5, 0.5))
}