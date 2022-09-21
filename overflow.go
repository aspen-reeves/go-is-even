package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const frac int = 32

func isEven(check int) bool {
	size := check * 2             // size of the array
	fracSize := int(size / frac)  // size of each fraction
	var array = make([]int, size) // array to hold the numbers
	var wg sync.WaitGroup         // wait group to wait for all the goroutines to finish
	wg.Add(frac)                  // add the number of goroutines to the wait group
	for i := 0; i < size; i++ {   // fill the array with zero
		array[i] = 0
	}

	for i := 0; i < frac; i++ {
		go func(i1 int) { // start a goroutine for each fraction
			defer wg.Done() // decrement the wait group when the goroutine is done
			for j := (fracSize * i1); j < (fracSize*i1 + fracSize); j++ {
				var a uint32 = uint32(j) // convert the index to a uint32
				for {

					a += 2      // add 2 to the number
					if a == 0 { // if the number overflows to 0, then it is even
						break
					}
					if a == 1 { // if the number overflows to 1, then it is odd
						break
					}
				}
				if a == 0 {
					array[j] = j // if the number is even, then set the number in the array to the index
				}
			}
		}(i)
	}
	wg.Wait()                                                                                  // wait for all the goroutines to finish
	var hmm bool = false                                                                       // variable to hold the result
	rand.Shuffle(len(array), func(i1, j1 int) { array[i1], array[j1] = array[j1], array[i1] }) // shuffle the array
	for x := 0; x < size; x++ {
		if array[x] == check { // if the number is in the array, then it is even
			hmm = true
			break
		}
	}
	return hmm // return the result

}

func main() {
	//a := 101 // change this number to test
	//fmt.Println("num to check:", a)
	//fmt.Scan(&a)
	//start := time.Now()
	for a := 100; a < 200; a++ {

		if isEvenRecursive(a, 0.0001) {
			fmt.Println(a, " is even")
		} else {
			fmt.Println(a, " is odd")
		}
	}

	//duration := time.Since(start)
	//fmt.Print("\ntime to complete:")
	//fmt.Println(duration)
	/*                           //uncomment this for prozper algo
	fmt.Println("properly done")
	start = time.Now()

	if isEvenNew(a) {
		fmt.Println("is even")
	} else {
		fmt.Println("is odd")
	}

	duration = time.Since(start)
	fmt.Print("\ntime to complete:")
	fmt.Println(duration)
	*/

}

//example output with ryzen 5 5700u
// Arraysize(maximum even check): 128

// number of fractions(sorta numbers of go routines): 32

// size of each array to check(arraysize/ fractions) 4

// num to check: 15

// 15 is not even
// 9.501672439s

//test code blocks to put in main
/*
fmt.Print("\nusing modulo\n")
	start := time.Now()
	if a%2 == 0 {
		fmt.Println(a, "is even")
	} else {
		fmt.Println(a, "is not even")
	}
	duration := time.Since(start)
	fmt.Print("\ntime to complete:")
	fmt.Println(duration)
*/
