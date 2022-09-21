package main

import (
	"math"
	"math/rand"
	"time"
)

func isEvenNew(check int) bool {
	temp := check
	even, odd := 0, 0
	for i := 0; i < check*100; i++ {
		rand.Seed((time.Now().UnixNano()))
		randInt := rand.Intn(temp)
		if randInt%2 == 0 {
			even++

		} else {
			odd++
		}
	}
	if (float64(even) / float64(odd)) > 1 {
		return false
	} else {
		return true
	}
}
func isItActuallyEven(check int) bool {
	evenProb := 0
	oddProb := 0
	conf := 1.3
	for j := 0; j < 10; j++ {
		if isEvenNew(check) {
			evenProb++
		} else {
			oddProb++
		}
	}
	for {
		if evenProb > int((conf*float64(oddProb))) || oddProb > int((conf*float64(evenProb))) {
			break
		} else {
			for j := 0; j < 10; j++ {
				if isEvenNew(check) {
					evenProb++
				} else {
					oddProb++
				}
			}
		}
	}
	if evenProb > oddProb {
		return true
	} else {
		return false
	}

}
func isEvenRecursive(check int) bool {
	temp := check
	even, odd := 0, 0
	for i := 0; i < check*100; i++ {
		rand.Seed((time.Now().UnixNano()))
		randInt := rand.Intn(temp)
		if randInt%2 == 0 {
			even++

		} else {
			odd++
		}
	}
	idealRatio := float64(check/2) / float64((check/2)+1)
	trueRatio := float64(even) / float64(odd)
	diff := math.Abs(float64(idealRatio - trueRatio))

}
