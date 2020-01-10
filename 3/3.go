package main

import "fmt"

func main() {
	a := 600851475143

	for i := 2; i < a; i++ {

		for a%i == 0 {

			a = a / i

		}

	}

	fmt.Println(a)
}
