package prime

import (
	"fmt"
	"math"
)

//IsPrime checks wether the passed number is prime
func IsPrime(number int) bool {

	if number < 2 {
		return false
	} else if number == 2 {
		return true
	}

	targetNumber := float64(number)
	modRealNumber := math.Mod(targetNumber, 2)

	if math.IsNaN(modRealNumber) {
		fmt.Println("Could not compute the rest")
	}

	modInteger, _ := math.Modf(modRealNumber)

	if modInteger == 0 {
		return false
	}

	for divisorToCheck := targetNumber - 1; divisorToCheck > 2; divisorToCheck-- {

		modRealNumber := math.Mod(targetNumber, divisorToCheck)
		if modRealNumber == 0 {
			return false
		}
	}

	return true

}
