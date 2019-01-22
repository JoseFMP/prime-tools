package testprimes

import (
	"testing"

	"github.com/JoseFMP/prime-tools/src/prime"
	"github.com/stretchr/testify/require"
)
  
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
	primesToTest[19] = true

	for num, isPrim := range primesToTest {
		require.Equalf(t, isPrim, prime.IsPrime(num), "Screwed up checking if %d is prime. Expected: %t", num, isPrim)
	}

}
