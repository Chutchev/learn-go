package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Printf("%T", x)
}

func adding(x, y string) int64 {
	xRunes := []rune(x)
	yRunes := []rune(y)
	digitXRunes := deleteSymbolsNotDigit(xRunes)
	digitYRunes := deleteSymbolsNotDigit(yRunes)
	//digitX, _ := strconv.ParseInt(digitXStr, 10, 64)
	//digitY, _ := strconv.ParseInt(digitYStr, 10, 64)
	fmt.Println(digitXRunes, digitYRunes)
	return 1
}

func deleteSymbolsNotDigit(r []rune) []rune {
	a := make([]rune, len(r), len(r))
	for i, el := range r {
		if unicode.IsDigit(el) {
			a[i] = el
		}
	}
	return a
}

func CSVDelit() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", ".")
	sArray := strings.Split(s, ";")
	a, _ := strconv.ParseFloat(sArray[0], 64)
	b, _ := strconv.ParseFloat(sArray[1], 64)
	fmt.Printf("%.4f", a/b)
}
