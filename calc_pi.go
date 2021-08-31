package main

import (
	"fmt"
	"math/rand"
	"flag"
	"github.com/joshuamhtsang/gophers_calc_pi/geometry"
)

func main() {
	fmt.Println("Gophers Calculate Pi!")

	// Parse arguments
	numIterationsPtr := flag.Int("n", 500, "Number of iterations.")
	flag.Parse()
	fmt.Println(*numIterationsPtr)

	fmt.Println(geometry.WithinUnitCircle(0.5, 0.5))

	rand.Seed(45)

	fmt.Println(estimate_pi((*numIterationsPtr)))
}

func estimate_pi(num_iterations int) float64 {
	var inside_quadrant = make([]bool, num_iterations)

	for i := 0; i < num_iterations; i++ {
		var x_test float64 = rand.Float64()
		var y_test float64 = rand.Float64()

		inside_quadrant[i] = geometry.WithinUnitCircle(x_test, y_test)
	}

	var fraction_inside float64 = compute_fraction_inside(inside_quadrant)

	return 4*fraction_inside
}

func compute_fraction_inside(inside []bool) float64 {
	var num_inside int = 0
	var num_samples int = len(inside)

	for i := 0; i < num_samples; i++ {
		if inside[i] {
			num_inside++
		}
	}

	var fraction float64 = float64(num_inside)/float64(num_samples)

	return fraction
}