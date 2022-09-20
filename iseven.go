package main

import (
	"fmt"
	"sync"
	"time"
)

const frac int = 32

func iseven(check int, arrSize int) {
	size := arrSize
	fracSize := int(size / frac)
	var array = make([]int, size)
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
	var hmm bool = false
	/*rand.Shuffle(len(array), func(i1, j1 int) { array[i1], array[j1] = array[j1], array[i1] })
	for x := 0; x < size; x++ {
		if array[x] == check {
			hmm = true
			break
		}
	}
	*/
	if array[check] == check {
		hmm = true
	}
	/*fmt.Println("even table")
	fmt.Println("number | even/odd")
	for q := 0; q < size; q++ {
		if array[q] == q {
			println(q, "    | even")
		} else {
			println(q, "    | odd")
		}
	}
	*/
	//fmt.Println()
	if hmm == true {
		fmt.Println(check, "is even")
	} else {
		fmt.Println(check, "is not even")
	}

}

func main() {
	a := 64
	b := a * 2
	fmt.Println("num to check:", a)
	fmt.Println()
	//test multiple array sizes
	/*for i := 1; i < 10; i++ {
		fmt.Println("computing with array size:", b*i)
		start := time.Now()
		iseven(a, b*i)
		duration := time.Since(start)
		fmt.Print("\ntime to complete:")
		fmt.Println(duration)
	}
	*/
	fmt.Println("computing with array size:", b)
	for {
		fmt.Println("my is-even function:")
		start := time.Now()
		iseven(a, b)
		duration := time.Since(start)
		fmt.Print("\ntime to complete:")
		fmt.Println(duration)
		break
	}
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
}

//example output with ryzen 5 5700u
// Arraysize(maximum even check): 128

// number of fractions(sorta numbers of go routines): 32

// size of each array to check(arraysize/ fractions) 4

// num to check: 15

// 15 is not even
// 9.501672439s
