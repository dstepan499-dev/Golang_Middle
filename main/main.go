package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	switch {
	case a < 10:
		for b > 0 {
			number := b % 10
			if number == a {
				fmt.Println(a)
				break
			}
			b /= 10
		}

	case a >= 10 && a < 100:
		for a > 0 {
			counter := 1
			number_a := a % 10

			number := b
			for number > 0 {
				number_b := number % 10

				if number_a == number_b {
					counter += 1
				}

				number = number / 10
			}

			a = a / 10
			if counter == 2 {
				fmt.Println(number_a)
			}
		}
	default:
		fmt.Println("Пошол НАХУЙ")
	}

	/*	for a > 0 {
		counter := 1
		number_a := a % 10

		number := b
		for number > 0 {
			number_b := number % 10

			if number_a == number_b {
				counter += 1
			}

			number = number / 10
		}

		a = a / 10
		if counter == 2 {
			fmt.Println(number_a)
		}
	}*/
}
