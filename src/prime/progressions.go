package prime

import (
	"fmt"
	"sync"
	"time"
)

func FindLengthArithmenticProgression(firstTerm int, difference int) int {

	length := 0
	for an := firstTerm; true; an += difference {
		if !IsPrime(an) {
			break
		}
		length++
	}
	return length
}

//FindLongestArithmeticProgression gives the maximum length and difference value of
//the arithmetic sequence of prime numbers whose initial term is passed as argument.
func FindLongestArithmeticProgression(initialTerm int, startingDifference int, waitTimeSecond int) (int, int) {

	startTimestamp := time.Now()

	maxLength := 0
	keyDifference := 0
	var ellapsedTime = float64(0)

	var routinesUsed = 0
	var routinesMax = 1000
	var lock sync.Mutex

	for difference := startingDifference; ellapsedTime < float64(waitTimeSecond); {

		if routinesUsed < routinesMax {
			routinesUsed++
			go func() {
				thisLength := FindLengthArithmenticProgression(initialTerm, difference)
				lock.Lock()
				if thisLength > maxLength {
					maxLength = thisLength
					keyDifference = difference
				}
				fmt.Printf("Current maximum length is %d, with difference %d \n", maxLength, difference)
				difference++
				routinesUsed--
				lock.Unlock()

			}()
		} else {
			time.Sleep(time.Second)
		}
		ellapsedTime = time.Now().Sub(startTimestamp).Seconds()
	}
	return maxLength, keyDifference

}
