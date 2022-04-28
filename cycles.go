package main

import (
	"fmt"
)

func main() {
	eigthCycle()
}

func first_cycle() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}
}

func second_cycle() {
	var firstNum, secondNum, summa int
	fmt.Scan(&firstNum)
	fmt.Scan(&secondNum)
	for i := firstNum; i <= secondNum; i++ {
		summa += i
	}
	fmt.Println(summa)
}

func third_cycle() {
	var i, j, summa int
	fmt.Scan(&i)
	for k := 0; k < i; k++ {
		fmt.Scan(&j)
		if j > 9 && j < 100 && j%8 == 0 {
			summa += j
		}
	}
	fmt.Println(summa)
}

func max_el_four_cycle() {
	var i, max, summa int
	summa = 1
	for fmt.Scan(&i); i != 0; fmt.Scan(&i) {
		if i > max {
			max = i
			summa = 1
		} else if i == max {
			summa++
		}
	}
	fmt.Println(summa)
}

func fifthCycle() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}

func sixthCycle() {
	var number int
	for fmt.Scan(&number); number < 100; fmt.Scan(&number) {
		if number < 10 {
			continue
		} else {
			fmt.Println(number)
		}
	}
}

func seventhCycle() {
	var x, p, y, i int
	fmt.Scan(&x, &p, &y)

	for i = 0; x < y; {
		x = x + (x*p)/100
		i++
	}
	fmt.Println(i)
}

func eigthCycle() {
	var firstNumber, secondNumber, ready string
	fmt.Scan(&firstNumber, &secondNumber)
	for i := 0; i < len(firstNumber); i++ {
		for j := 0; j < len(secondNumber); j++ {
			if firstNumber[i] == secondNumber[j] {
				ready += fmt.Sprintf("%v ", string(firstNumber[i]))
			}
		}
	}
	fmt.Println(ready[:len(ready)])
}
