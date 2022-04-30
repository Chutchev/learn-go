package main

import "fmt"

func main() {
	cifr_kor()
}

func first_task_one() {
	var a, summa int
	fmt.Scan(&a)
	summa += a / 100
	summa += a / 10 % 10
	summa += a % 10
	fmt.Println(summa)

}

func second_task_one() {
	var a int
	fmt.Scan(&a)
	fmt.Print(a%10, "")
	fmt.Print(a/10%10, "")
	fmt.Print(a/100, "")
}

func how_time_from_seconds() {
	var seconds int
	fmt.Scan(&seconds)
	hour := seconds / 3600
	minutes := (seconds - hour*3600) / 60
	fmt.Printf("It is %d hours %d minutes.", hour, minutes)
}

func what_triangle() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}

func is_triangle_ex() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a+b > c && b+c > a && c+a > b {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}

func middle() {
	var a, b int
	fmt.Scan(&a, &b)
	c := float64(a+b) / 2.0
	fmt.Println(c)
}

func zero_count() {
	var count, number, a int
	fmt.Scan(&number)
	for i := 0; i < number; i++ {
		fmt.Scan(&a)
		if a == 0 {
			count++
		}
	}
	fmt.Println(count)
}

func count_of_min_value() {
	var count, number int
	fmt.Scan(&number)
	var array = make([]int, number, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&array[i])
	}
	min := array[0]
	for _, el := range array {
		if el < min {
			min = el
			count = 1
		} else if el == min {
			count++
		}
	}
	fmt.Println(count)
}

func cifr_kor() {
	var summa, number int
	fmt.Scan(&number)
	for number >= 10 {
		summa += number % 10
		number = number / 10
	}

	if summa >= 10 {
		summa = summa/10 + summa%10
	}
	fmt.Println(summa)
	fmt.Println(summa)
}
