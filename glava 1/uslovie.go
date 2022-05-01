package main

import "fmt"

func main() {
	year_is_vis()
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

func second_usl() {
	var number int
	fmt.Scan(&number)
	f := number / 100
	s := number % 100 / 10
	l := number % 10
	if f != s && f != l && s != l {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func third_usl() {
	var number, f int
	fmt.Scan(&number)
	switch {
	case 0 < number && number < 9:
		fmt.Println(number)
	case 9 < number && number < 100:
		f = number / 10
		fmt.Println(f)
	case 100 <= number && number < 1000:
		f = number / 100
		fmt.Println(f)
	case 1000 <= number && number < 10000:
		f = number / 1000
		fmt.Println(f)
	case number == 10000:
		f = 1
		fmt.Println(f)
	}
}

func fourth_usl() {
	var number int
	fmt.Scan(&number)
	f := int(number / 100000)
	s := int(number / 10000 % 10)
	t := int(number / 1000 % 10)
	fo := int(number % 1000 / 100)
	fi := int(number % 100 / 10)
	six := int(number % 10)
	sum1 := sum(f, s, t)
	sum2 := sum(fo, fi, six)
	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func sum(f int, s int, t int) int {
	return f + s + t
}

func year_is_vis() {
	var year int
	fmt.Scan(&year)
	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
