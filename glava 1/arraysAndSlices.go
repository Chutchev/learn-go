package main

import "fmt"

// make([]T, length, capacity) <- создать срез

func main() {
	fourthSlice()
}

func firstSlice() {
	var length int
	fmt.Scan(&length)
	a := make([]int, length, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(a[3])
}

func secondSlice() {
	array := [5]int{}
	var a, max int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	max = array[0]
	for _, el := range array {
		if el > max {
			max = el
		}
	}
	fmt.Println(max)
}

func thirdSlice() {
	var size int
	fmt.Scan(&size)
	slice := make([]int, size, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&slice[i])
	}
	for index, el := range slice {
		if index%2 == 0 {
			fmt.Print(el, " ")
		}
	}
}

func fourthSlice() {
	var size, summa int
	fmt.Scan(&size)
	slice := make([]int, size, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&slice[i])
	}
	for _, el := range slice {
		if el > 0 {
			summa += 1
		}
	}
	fmt.Println(summa)
}
