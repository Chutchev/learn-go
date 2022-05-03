package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Average struct {
	Average float32
}

func main() {
	//file, err := os.Open("./glava 3/test.json")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer file.Close()
	var (
		group   Group
		average Average
	)
	file, err := ioutil.ReadAll(os.Stdin)
	//s, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &group)
	//err = json.Unmarshal(s, &group)
	if err != nil {
		return
	}
	for _, student := range group.Students {
		average.Average += float32(len(student.Rating))
	}
	average.Average = average.Average / float32(len(group.Students))
	info, errJson := json.MarshalIndent(average, "", "    ")
	if errJson != nil {
		fmt.Println(err)
		return
	}
	io.WriteString(os.Stdout, string(info))
}
