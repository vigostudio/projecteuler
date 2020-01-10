package main

import "fmt"

func main() {
	m := 20
	for i := 20; i > 2; i-- {
		if m%i != 0 {
			m = LCM(m, i)
		}
		fmt.Println(i)
		fmt.Println(m)
	}
	fmt.Println(m)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {

	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}
