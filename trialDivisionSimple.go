package main

import "math"

func trialDivisionSimple(n uint16) bool {
	root_n := uint16(math.Sqrt(float64(n)))

	var i uint16
	for i = 3; i < root_n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
