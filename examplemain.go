package main

import "fmt"

func main() {

	for i := 0; i < 1000; i++ {
		if isEvenRoot(i) != (i%2 == 0) {
			fmt.Println(i)
		}
	}

}
