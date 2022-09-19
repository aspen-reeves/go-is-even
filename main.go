package main

import (
	"fmt"
	"sync"
)

const size int = 32000
const frac int = 16
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
	iseven(13)

}
