package main

import (
	"math"
	"math/rand"
	"time"
)

func isEvenNew(check int) bool {
	temp := check
	even, odd := 0, 0
	for i := 0; i < check*100; i++ { //generates check*100 random numbers. if ratio between even/odd is x/x it is even, x/x+1 if odd
		rand.Seed((time.Now().UnixNano()))
		randInt := rand.Intn(temp) // generate random
		if randInt%2 == 0 {        // checks if random num is even or odd
			even++ //iterates even counter if even

		} else {
			odd++ // iterates odd counter if odd
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
func isEvenRecursive(check int, confidence float64) bool { // unfinished
	temp := check
	even, odd := 0, 0
	for i := 0; i < check*100; i++ {
		rand.Seed((time.Now().UnixNano()))
		randInt := rand.Intn(temp) + 1 //generate random numbers from 1:
		if randInt%2 == 0 {
			even++

		} else if randInt%2 == 1 {
			odd++
		}
	}
	oddRatio := float64(check/2) / float64((check/2)+1) // ideal ratio if the number is odd
	evenRatio := 1.0                                    // ideal ratio if number is even

	trueRatio := float64(even) / float64(odd) // calcuates the true ratio of even/odd
	//fmt.Println(trueRatio, oddRatio, evenRatio)

	oddDiff := math.Abs(float64(oddRatio - trueRatio))   // difference between ideal and true ratios(odd)
	evenDiff := math.Abs(float64(evenRatio - trueRatio)) // difference between ideal and true ratios(even)

	if oddDiff < evenDiff { // if the odd diff is higher, check how close
		if trueRatio < oddRatio+(oddRatio*confidence) && trueRatio > oddRatio-(oddRatio*confidence) {
			return false
		} else {
			return isEvenRecursive(check, confidence) // recursive if it is too far off ideal
		}
	} else {
		if trueRatio < evenRatio+(evenRatio*confidence) && trueRatio > evenRatio-(evenRatio*confidence) {
			return true
		} else {
			return isEvenRecursive(check, confidence)
		}
	}

}
