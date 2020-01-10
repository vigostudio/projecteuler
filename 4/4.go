package main

import "fmt"

func main() {

	s := 999
	t := 999

	for j := 999; j > 900; j-- {

		s = j

		for k := j; k > 900; k-- {

			t = k

			if check(k * j) {
				break
			}

		}

		if check(t * j) {
			break
		}

	}

	fmt.Println(s * t)
	fmt.Println(s)
	fmt.Println(t)

	fmt.Println(check(10011))
}

func check(a int) bool {

	number := a

	var remainder int
	var reverse int = 0

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10
		if number == 0 {
			break
		}
	}

	return reverse == a

}
