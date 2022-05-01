package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	goodPassword()
}

func rightStrokes() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")
	textRunes := []rune(text)
	if unicode.IsUpper(textRunes[0]) && string(textRunes[len(textRunes)-2]) == "." {
		fmt.Println(text)
	}
}

func palindrome() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, " ", "")
	textRunes := []rune(text)
	reversedTextRunes := make([]rune, len(textRunes), len(textRunes))
	for i := len(textRunes) - 1; i >= 0; i-- {
		reversedTextRunes = append(reversedTextRunes, textRunes[i])
	}
	reversedTextRunes = reversedTextRunes[len(textRunes):]
	reversedText := string(reversedTextRunes)

	if reversedText == text {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}

func everySecondChar() {
	var a string
	fmt.Scan(&a)
	aRune := []rune(a)
	for i, el := range aRune {
		if i%2 == 1 {
			fmt.Print(string(el))
		}
	}
}

func counter() {
	var a string
	fmt.Scan(&a)
	aRune := []rune(a)
	//aNewRune := make([]rune, len(a), len(a))
	for _, el := range aRune {
		count := strings.Count(a, string(el))
		if count == 1 {
			fmt.Print(string(el))
		}
	}
}

func goodPassword() {
	var (
		password      string
		rightPassword bool
	)
	fmt.Scan(&password)
	passwordRune := []rune(password)
	if len(passwordRune) >= 5 {
		for _, el := range passwordRune {
			if !unicode.Is(unicode.Latin, el) && !unicode.IsDigit(el) {
				rightPassword = false
				fmt.Println("Wrong password")
				break
			} else if unicode.IsPunct(el) {
				rightPassword = false
				fmt.Println("Wrong password")
				break
			} else {
				rightPassword = true
			}
		}
	} else {
		fmt.Println("Wrong password")
	}
	if rightPassword {
		fmt.Println("Ok")
	}

}
