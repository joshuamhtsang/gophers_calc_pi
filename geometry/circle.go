package geometry

import "math"

func WithinUnitCircle(x float64, y float64) bool {
	var a float64 = math.Sqrt(x*x + y*y)
	
	if a < 1.0 {
		return true
	} else {
		return false
	}
}
