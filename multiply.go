package main

import (
	"fmt"
	"math"
)

// 1.5 stepik
func main() {
	count_of_hours()
}

func first() {
	var a int
	fmt.Scan(&a)
	a *= 2
	a += 100
	fmt.Println(a)
}

func second() {
	var a, b, c int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	a = a * a
	b = b * b
	c = a + b
	fmt.Println(c)
}

func square() {
	var a int
	fmt.Scan(&a)
	a = int(math.Pow(float64(a), 2))
	fmt.Println(a)
}

func last_number() {
	var a int
	fmt.Scan(&a)
	a = a % 10
	fmt.Println(a)
}

func count_of_dozens() {
	var a int
	fmt.Scan(&a)
	a = a % 100 / 10
	fmt.Println(a)
}

func count_of_hours() {
	var d, hours, minutes int
	fmt.Scan(&d)
	hours = d / 30
	minutes = d % 30
	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}
