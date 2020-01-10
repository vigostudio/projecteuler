package main

import (
	"fmt"
	"math"
)

func main() {

	prime := 1

	primeIndex := 1

	for i := 2; i <= 8; i++ {

		if check(i) {

			prime = i
			primeIndex++

		}
		fmt.Println(prime)
	}

	fmt.Println(prime)
	fmt.Println(primeIndex)
}

func check(n int) bool {
	i := 2
	for ; float64(i) < math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			break
		}
	}

	return i == n

}
