package main

import (
	"fmt"
	"io"
)

func main() {
	var n int
	var maxNumber int
	counter := 0

	for {
		_, err := fmt.Scan(&n)
		
		if err != nil {
			if err == io.EOF {
				// Достигнут конец ввода (EOF)
				break
			}

			fmt.Println("Ошибка ввода")
			return
		}

		// Ноль (0) - признак окончания последовательности
		if n == 0 {
			break
		}

		if counter == 0 || n > maxNumber {
			maxNumber = n
			counter = 1
		} else if n == maxNumber {
			counter++
		}
	}

	fmt.Println(counter)
} 