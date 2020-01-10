package main

import "fmt"

var ar []int

func main() {

	ar = make([]int, 50)

	sum := 0

	//ar := []int{}

	ar[0] = 1
	ar[1] = 1

	for i := 2; i < 50 && ar[i-1] < 4000000; i++ {

		ar[i] = ar[i-2] + ar[i-1]
		if ar[i]%2 == 0 {
			sum = sum + ar[i]
		}
	}

	fmt.Println(sum)
}

func fib(a int) int {

	if a == 1 || a == 0 {
		return 1
	}

	if a < 4000000 {
		return fib(a-1) + fib(a-2)
	}

	return fib(a - 1)

}
