package main

import "fmt"

func main() {
	fmt.Println(sumInt(1, 0))
}

func minimumFromFour() int {
	var a, min int
	fmt.Scan(&a)
	min = a
	for i := 0; i < 3; i++ {
		fmt.Scan(&a)
		if a < min {
			min = a
		}
	}
	return min
}

func vote(x int, y int, z int) int {
	if x == y || x == z {
		return x
	} else {
		return y
	}
}

func fibonacci(n int) int {
	table := make([]int, n+1)
	table[0] = 0
	table[1] = 1

	for i := 2; i <= n; i += 1 {

		table[i] = table[i-1] + table[i-2]

	}

	return table[n]
}

func sumInt(a ...int) (int, int) {
	summa := 0
	for _, el := range a {
		summa += el
	}
	return len(a), summa
}
