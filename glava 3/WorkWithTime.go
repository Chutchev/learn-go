package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	getDuration()
}

func toUnixDate() {
	var dataTime string
	fmt.Scan(&dataTime)
	firstTime, err := time.Parse(time.RFC3339, dataTime)
	if err != nil {
		panic(err)
	}
	fmt.Println(firstTime.Format(time.UnixDate))
}

func nextDay() {
	var dataTime string
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	dataTime = scan.Text()
	firstTime, err := time.Parse("2006-01-02 15:04:05", dataTime)
	if err != nil {
		panic(err)
	}
	fmt.Println(firstTime.Hour())
	if firstTime.Hour() > 12 {
		firstTime = firstTime.AddDate(0, 0, 1)
	}
	fmt.Println(firstTime.Format("2006-01-02 15:04:05"))
}

func getDuration() {
	var dataTimes string
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	dataTimes = scan.Text()
	dates := strings.Split(dataTimes, ",")
	date1Str, date2Str := dates[0], dates[1]
	date1, _ := time.Parse("02.01.2006 15:04:05", date1Str)
	date2, _ := time.Parse("02.01.2006 15:04:05", date2Str)
	if date1.After(date2) {
		fmt.Println(date1.Sub(date2))
	} else {
		fmt.Println(date2.Sub(date1))
	}
	fmt.Println(date1, date2)
}
