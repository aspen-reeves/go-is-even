package main

import "fmt"

func main() {
	//a := 101 // change this number to test
	//fmt.Println("num to check:", a)
	//fmt.Scan(&a)
	//start := time.Now()
	isRight := 0
	isWrong := 0
	for a := 0; a < 2000; {

		if isEvenRecursive(a, 0.00001) {
			fmt.Println(a, " is even")
			if a%2 == 0 {
				isRight++
			} else {
				isWrong++
			}
		} else {
			fmt.Println(a, " is odd")
			if a%2 == 1 {
				isRight++
			} else {
				isWrong++
			}
		}
		acc := float64(isRight) / float64(isRight+isWrong)
		acc *= 100
		fmt.Println("Accuracy:", acc, "%")

	}
}
