package main

import (
	"errors"
	"fmt"
)

func main() {
	var a uint32 = 5
	defer SpecPrint(a)
	a = 6
	fmt.Print(a)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		fmt.Println("Что то пошло не так!")
		return -1, errors.New("Деление на 0")
	} else {
		fmt.Println("Тут все норм")
		return a / b, nil
	}
}

func SpecPrint(s uint32) uint32 {
	fmt.Print(s)
	s++
	return s
}
