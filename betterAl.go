package main

import (
	"fmt"
	"math/rand"
	"time"
)

func isEvenNew(check int) bool {
	var prob float64
	temp := check
	var a, b float64
	for i := 0; i < temp*100; i++ {
		rand.Seed(time.Now().UnixNano())
		randInt := rand.Intn(temp)
		if randInt%2 == 0 {
			a++

		} else {
			b++
		}
	}
	prob = a / b
	fmt.Println(prob)
	if prob >= 1 {
		return false
	} else {
		return true
	}
}
func isItActuallyEven(check int) bool {
	evenProb := 0
	oddProb := 0
	for i := 0; i < 1000; i++ {
		if isEvenNew(check) {
			evenProb++
		} else {
			oddProb++
		}
	}
	if evenProb > oddProb {
		return true
	} else {
		return false
	}

}
