package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var power string
	fmt.Scan(&power)
	b := Battary{power: power}
	fmt.Println(b)
}

func readTask() (value1, value2, operation interface{}) {

	return 5.0, 5.6, "/" //тут играемся с параметрами

}

func sumFloatsInterfaces() {
	value1, value2, operation := readTask()
	switch v := value1.(type) {
	case float64:
		break
	default:
		fmt.Printf("value=%v: %T", v, v)
		os.Exit(0)
	}
	switch v := value2.(type) {
	case float64:
		break
	default:
		fmt.Printf("value=%v: %T", v, v)
		os.Exit(0)
	}
	switch operation.(type) {
	case string:
		break
	default:
		fmt.Println("неизвестная операция")
		os.Exit(0)
	}
	switch operation {
	case "+":
		res := value1.(float64) + value2.(float64)
		fmt.Printf("%.4f", res)
	case "-":
		res := value1.(float64) - value2.(float64)
		fmt.Printf("%.4f", res)
	case "*":
		res := value1.(float64) * value2.(float64)
		fmt.Printf("%.4f", res)
	case "/":
		res := value1.(float64) / value2.(float64)
		fmt.Printf("%.4f", res)
	}
}

type Battary struct {
	power string
}

func (b Battary) String() string {
	battaryPower := ""
	//countX := strings.Count(b.power, "1")
	countZ := strings.Count(b.power, "0")
	for i := 0; i < 10; i++ {
		if i < countZ {
			battaryPower += " "
		} else {
			battaryPower += "X"
		}
	}
	return fmt.Sprintf("[%s]", battaryPower)
}
