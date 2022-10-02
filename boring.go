package main

func isEven(check int) bool {
	if check%2 == 0 {
		return true
	} else {
		return isOdd(check)
	}
}
func isOdd(check int) bool {
	if check%2 == 1 {
		return true
	} else {
		return isEven(check)
	}
}

func bitWiseEven(check int) bool {
	return (check&1 == 0)
}
