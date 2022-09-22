package main

import (
	"math/rand"
	"sync"
)

func isEvenOverflow(check int) bool {
	const frac int = 32
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
