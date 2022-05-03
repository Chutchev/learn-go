package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	findZero()
}
func workWithBuf() {
	var summa = 0
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		num, _ := strconv.Atoi(scan.Text())
		summa += num
	}
	readySumma := strconv.Itoa(summa)
	io.WriteString(os.Stdout, readySumma)
}
func findZero() {
	f, err_open := os.Open("./glava 3/task.data")
	defer f.Close()
	if err_open != nil {
		fmt.Println(err_open)
	}
	r := csv.NewReader(f)
	data, err_read := r.ReadAll()
	if err_read != nil {
		fmt.Println(err_open)
	}
	array := strings.Split(data[0][0], ";")
	for index, el := range array {
		if el == "0" {
			fmt.Println(index + 1)
			break
		}
	}
}
