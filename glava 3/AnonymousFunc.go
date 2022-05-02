package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(a uint) uint {
		var ready = ""
		b := strconv.Itoa(int(a))
		for _, el := range b {
			v, _ := strconv.Atoi(string(el))
			if v%2 == 0 && v != 0 {
				ready += strconv.Itoa(v)
			}
		}
		n, _ := strconv.Atoi(ready)
		if n == 0 {
			return 100
		}
		return uint(n)
	}
	fmt.Println(fn(0))
}
