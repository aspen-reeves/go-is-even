package main

import (
	"fmt"
	"sync"
	"time"
)

const size int = 128
const frac int = 32
const fracSize int = int(size / frac)

func iseven(check int) {
	var array [size]int
	var wg sync.WaitGroup
	wg.Add(frac)
	for i := 0; i < size; i++ {
		array[i] = 0
	}

	for i := 0; i < frac; i++ {
		go func(i1 int) {
			defer wg.Done()
			for j := (fracSize * i1); j < (fracSize*i1 + fracSize); j++ {
				var a uint32 = uint32(j)
				for {
					a++
					a++
					if a == 0 {
						break
					}
					if a == 1 {
						break
					}
				}
				if a == 0 {
					array[j] = j
				}
			}
		}(i)
	}
	wg.Wait()
	if array[check] != 0 {
		fmt.Println(check, "is even")
	} else {
		fmt.Println(check, "is not even")
	}

}

func main() {
	fmt.Println("Arraysize(maximum even check):", size)
	fmt.Println()
	fmt.Println("number of fractions(sorta numbers of go routines):", frac)
	fmt.Println()
	fmt.Println("size of each array to check(arraysize/ fractions)", fracSize)
	fmt.Println()
	a := 15
	fmt.Println("num to check:", a)
	fmt.Println()
	start := time.Now()
	iseven(a)
	duration := time.Since(start)
	fmt.Println(duration)

}
