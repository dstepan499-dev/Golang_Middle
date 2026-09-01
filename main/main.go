package main

import "fmt"

func main() {
	var n int

	for {
		fmt.Scan(&n)

		if n > 100 {
			break
		} else if n < 10 {
			continue
		} else {
			fmt.Println(n)
		}
	}
}
