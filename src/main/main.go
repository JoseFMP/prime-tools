package main

import "fmt"
import ".."

func main() {
	maxLength, diff := prime.FindLongestArithmeticProgression(5, 253927, 5000)

	fmt.Printf("Max length found: %d, difference: %d \n", maxLength, diff)
}
