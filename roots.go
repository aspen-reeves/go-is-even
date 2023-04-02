package main

import (
	"math"
	"math/rand"
	"time"
)

func isEvenRoot(check int) bool {
	// we will imagine that the check number is the degree of a polynomial
	// and we will check if the polynomial is even or odd, using the roots of the polynomial

	//generate random number
	const confidence = 100000000
	rand.Seed((time.Now().UnixNano()))
	randInt := rand.Intn(1000) - 500
	x := float64(randInt)

	fx := (math.Pow(x, float64(check)) + 1)
	dfx := float64(check) * math.Pow(x, float64(check-1))
	x1 := x - (fx / dfx)

	for i := 0; i < confidence; i++ {
		x = x1
		fx = (math.Pow(x, float64(check)) + 1)
		dfx = float64(check) * math.Pow(x, float64(check-1))
		x1 = x - (fx / dfx)
		if math.Abs(x1-x) < 1.0/confidence {
			return false
		}
	}
	return true
}
