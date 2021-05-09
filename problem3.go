package euler

import (
	"math"
)

/*
Problem  #3
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

var maxq float64

func Problem3(num, divisor float64) float64 {

	var quotient = num

	if math.Mod(quotient, divisor) == 0 {
		quotient = num / divisor
		maxq = quotient
	} else {
		divisor += 1
	}

	if quotient > 0 && divisor < math.Sqrt(quotient) {
		Problem3(quotient, divisor)
	}

	return maxq
}
