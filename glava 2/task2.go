package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var k, p, v float64 = 1296, 6, 6

func main() {
	T()
}

func hypot() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(math.Sqrt(float64(a*a + b*b)))
}

func star() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\r\n")
	for index, el := range text {
		//fmt.Print(index, string(el), "\n")
		if index == len(text)-1 {
			fmt.Print(string(el))
		} else {
			fmt.Print(string(el), "*")
		}

	}
}

func maxValue() {
	var text string
	fmt.Scan(&text)
	var max int32 = int32(text[0])
	for _, el := range text {
		if el > max {
			max = el
		}
	}
	fmt.Println(string(max))
}

func powNumber() {
	var text string
	fmt.Scan(&text)
	for _, el := range text {
		fmt.Print((el - '0') * (el - '0'))
	}
}

func T() float64 {
	return 6 / W()
}

func W() float64 {
	return math.Sqrt(float64(k / M()))
}

func M() float64 {
	return float64(p * v)
}
