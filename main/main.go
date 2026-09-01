package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)

	var total float64 = float64(x)
	year := 0

	for int(total) < y {
		year += 1
		var percent float64 = total / 100 * float64(p)
		total += percent
	}

	fmt.Println(year)
}
