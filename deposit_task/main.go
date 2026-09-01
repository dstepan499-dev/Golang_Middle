package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)

	// Переводим все в копейки
	x *= 100
	y *= 100

	years := 0

	for x < y {
		// Целочисленное деление само отбросит доли копеек
		x += (x * p) / 100
		years++
	}

	fmt.Println(years)
}
