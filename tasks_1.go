package main

import (
	"fmt"
	"math"
)

func main() {
	deleteNumber()
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
	for number > 0 {
		summa += number % 10
		number = number / 10
	}
	fmt.Println(summa)
	if summa >= 10 {
		summa = summa/10 + summa%10
	}
	fmt.Println(summa)
}

func biggerHundred() {
	var a, b, c int
	var isExists bool = false
	fmt.Scan(&a, &b)
	for i := b; i >= a; i-- {
		if i%7 == 0 {
			c = i
			isExists = true
			break
		}
	}
	if isExists {
		fmt.Println(c)
	} else {
		fmt.Println("NO")
	}
}

func howManyCow() {
	var cow int
	var stroke string
	fmt.Scan(&cow)
	lastCow := cow
	if cow > 20 {
		lastCow = cow % 10
	}
	switch {
	case lastCow == 1:
		stroke = "korova"
	case lastCow >= 2 && lastCow <= 4:
		stroke = "korovy"
	case lastCow > 4 && lastCow <= 20:
		stroke = "korov"

	}
	fmt.Printf("%d %s", cow, stroke)
}

func pow() {
	var N uint16
	fmt.Scan(&N)
	for i := 0; ; i++ {
		a := math.Pow(2, float64(i))
		if uint16(a) <= N {
			fmt.Print(a, " ")
		} else {
			break
		}
	}
}

func countOfFib() {
	var A int
	fmt.Scan(&A)
	i := numberOfFib(0, 1, 2, A)
	fmt.Println(i)
}

func numberOfFib(a int, b int, i int, N int) int {
	nextFib := a + b
	if nextFib == N {
		return i
	} else if nextFib > N {
		return -1
	} else {
		return numberOfFib(b, nextFib, i+1, N)
	}
}

func dual() {
	var N int
	fmt.Scan(&N)
	fmt.Printf("%b", N)
}

func deleteNumber() {
	var number, digit int
	fmt.Scan(&number, &digit)
	var array = make([]int, 0, len(fmt.Sprintf("%v", number)))
	for number > 0 {
		fmt.Println(number, array)
		if number%10 != digit {
			array = append(array, number%10)
		}
		number = number / 10
	}
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Print(array[i])
	}
}
