package main

import "fmt"

func main() {
	first_usl()
}

func first_usl() {
	var c int
	fmt.Scan(&c)
	switch {
	case 0 > c:
		fmt.Println("Отрицательное число")
	case 0 == c:
		fmt.Println("Ноль")
	case c > 0:
		fmt.Println("Положительное число")
	}
}
