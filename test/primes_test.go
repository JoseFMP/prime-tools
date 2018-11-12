package testprimes

import "testing"
import "../src"

func TestPrimes(t *testing.T) {

	primesToTest := make(map[int]bool)

	primesToTest[0] = false
	primesToTest[1] = false
	primesToTest[2] = true
	primesToTest[3] = true
	primesToTest[4] = false
	primesToTest[5] = true
	primesToTest[6] = false
	primesToTest[7] = true
	primesToTest[8] = false
	primesToTest[9] = false
	primesToTest[10] = false
	primesToTest[11] = true
	primesToTest[12] = false
	primesToTest[13] = true
	primesToTest[14] = false
	primesToTest[15] = false
	primesToTest[16] = false
	primesToTest[17] = true
	primesToTest[18] = false

	for num, isPrim := range primesToTest {

		if prime.IsPrime(num) != isPrim {
			t.Errorf("Screwed up checking if %d is prime. Expected: %t", num, isPrim)
			t.FailNow()
		}
	}

}
