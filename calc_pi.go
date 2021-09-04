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

	fmt.Println(estimatePi((*numIterationsPtr)))

	// Calculate with go routines.
	estimate_10 := make(chan float64)
	estimate_100 := make(chan float64)
	estimate_1000 := make(chan float64)
	estimate_10000 := make(chan float64)
	estimate_10000000 := make(chan float64)
	go goEstimatePi(10, estimate_10)
	go goEstimatePi(100, estimate_100)
	go goEstimatePi(1000, estimate_1000)
	go goEstimatePi(10000, estimate_10000)
	go goEstimatePi(10000000, estimate_10000000)
	fmt.Println("estimates from 10000000: ", <-estimate_10000000)
	fmt.Println("estimates from 10: ", <-estimate_10)
	fmt.Println("estimates from 100: ", <-estimate_100)
	fmt.Println("estimates from 1000: ", <-estimate_1000)
	fmt.Println("estimates from 10000: ", <-estimate_10000)
	
}

func goEstimatePi(num_iterations int, estimates chan float64) {
	estimate := estimatePi(num_iterations)
	estimates <- estimate
}

func estimatePi(num_iterations int) float64 {
	var inside_quadrant = make([]bool, num_iterations)

	for i := 0; i < num_iterations; i++ {
		var x_test float64 = rand.Float64()
		var y_test float64 = rand.Float64()

		inside_quadrant[i] = geometry.WithinUnitCircle(x_test, y_test)
	}

	var fraction_inside float64 = computeFractionInside(inside_quadrant)

	return 4*fraction_inside
}

func computeFractionInside(inside []bool) float64 {
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