package main

import (
	"fmt"
	"time"
)

/*
	func main() {
		//use bitwise even meathod
		start := time.Now()
		num := 1000000000001
		check := (num&1 == 0) //bitWiseEven(num)
		duration := time.Since(start)

		fmt.Println("Time taken:", duration)
		if check {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
		start = time.Now()
		check = (num%2 == 0)
		duration = time.Since(start)

		fmt.Println("Time taken:", duration)
		if check {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
*/
func main() {
	fmt.Println("Welcome to the even number checker!")
	fmt.Print("Use custom logic? (y/n)")
	var custom string
	fmt.Scanln(&custom)
	if custom == "y" {
		fmt.Println("no")
	} else {
		fmt.Println("Using default logic.")

		fmt.Print("\nPlease enter a number to check if it is even or odd:")
		var check int
		fmt.Scanln(&check)

		fmt.Println("Select method for checking if the number is even:")
		fmt.Println("0. Standard(modulo) method")
		fmt.Println("1. Overflow method")
		fmt.Println("2. Probability method")
		fmt.Println("3. Double-check probability method")
		fmt.Println("4. Recursive probability method")
		var method int
		fmt.Scanln(&method)

		fmt.Print("Do you want benchmarking? (y/n)")
		var bench string
		fmt.Scanln(&bench)

		var isBenchmarking bool
		if bench == "y" {
			isBenchmarking = true
		} else {
			isBenchmarking = false
		}
		start := time.Now()
		switch method {
		case 0:
			fmt.Println("Standard method selected")
			fmt.Println("Checking if", check, "is even...")
			if isEven(check) {
				fmt.Println(check, "is even!")
			} else {
				fmt.Println(check, "is odd!")
			}

		case 1:
			fmt.Println("Overflow method selected")
			fmt.Println("Checking if", check, "is even...")
			if isEvenOverflow(check) {
				fmt.Println(check, "is even!")
			} else {
				fmt.Println(check, "is odd!")
			}
		case 2:
			fmt.Println("Probability method selected")
			fmt.Println("Checking if", check, "is even...")
			if isEvenProb(check) {
				fmt.Println(check, "is even!")
			} else {
				fmt.Println(check, "is odd!")
			}
		case 3:
			fmt.Println("Double-check probability method selected")
			fmt.Println("Checking if", check, "is even...")
			if isItActuallyEven(check) {
				fmt.Println(check, "is even!")
			} else {
				fmt.Println(check, "is odd!")
			}
		case 4:
			fmt.Println("Recursive method selected")
			fmt.Println("Enter confidence level (0.0001 is default)")
			var confidence float64
			fmt.Scanln(&confidence)

			fmt.Println("Checking if", check, "is even...")
			if isEvenRecursive(check, confidence) {
				fmt.Println(check, "is even!")
			} else {
				fmt.Println(check, "is odd!")
			}
		}
		duration := time.Since(start)
		if isBenchmarking {
			fmt.Println("Time taken:", duration)
		}

	}
}

/*func main() {
	//start := time.Now()
	isRight := 0
	isWrong := 0
	for a := 2; a < 2000; a++ {

		if isEvenRecursive(a, 0.0002) {
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
*/
