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
func isEvenRecursive(check int, even int, odd int, confidence float64) bool { // unfinished
	temp := check
	for i := 0; i < check*100; i++ {
		rand.Seed((time.Now().UnixNano()))
		randInt := rand.Intn(temp)
		if randInt%2 == 0 {
			even++

		} else {
			odd++
		}
	}
	oddRatio := float64(check/2) / float64((check/2)+1) // ideal ratio if the number is odd
	evenRatio := 1.0

	trueRatio := float64(even) / float64(odd)

	oddDiff := math.Abs(float64(oddRatio - trueRatio))
	evenDiff := math.Abs(float64(evenRatio - trueRatio))

	if oddDiff < evenDiff {
		if trueRatio < oddRatio+(oddRatio*confidence) && trueRatio > oddRatio-(oddRatio*confidence) {
			return false
		} else {
			return isEvenRecursive(check, even, odd, confidence)
		}
	} else {
		if trueRatio < evenRatio+(evenRatio*confidence) && trueRatio > evenRatio-(evenRatio*confidence) {
			return true
		} else {
			return isEvenRecursive(check, even, odd, confidence)
		}
	}

}
