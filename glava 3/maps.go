package main

//
//import (
//	"fmt"
//)
//
//func main() {
//	groupCity := map[int][]string{
//		10:   []string{"Деревня", "Село"},        // города с населением 10-99 тыс. человек
//		100:  []string{"Город", "Большой город"}, // города с населением 100-999 тыс. человек
//		1000: []string{"Миллионик"},              // города с населением 1000 тыс. человек и более
//	}
//	cityPopulation := map[string]int{
//		"Село":      100,
//		"Миллионик": 500,
//	}
//}
//
//func work(n int) {
//	var a int
//	dictionary := make(map[int]int)
//	for i := 0; i < 10; i++ {
//		fmt.Scan(&a)
//		if value, ok := dictionary[a]; ok {
//			fmt.Print(value, " ")
//		} else {
//			//dictionary[a] = work(a)
//			fmt.Print(dictionary[a], " ")
//		}
//	}
//}
