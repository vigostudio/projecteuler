package main

import (
	"fmt"
	"math"
)

func main() {

	for a := 1; a <= 1000; a++ {

		for b := a + 1; b <= 1000; b++ {

			c := (a * a) + (b * b)

			if math.Sqrt(float64(c))-math.Floor(math.Sqrt(float64(c))) == 0 {
				if (float64(a) + float64(b) + math.Sqrt(float64(c))) == 1000 {

					fmt.Println(a)
					fmt.Println(b)
					fmt.Println(c)
					fmt.Println(float64(a) * float64(b) * math.Sqrt(float64(c)))

				}

				//fmt.Println("--")

			}

		}

	}

}
