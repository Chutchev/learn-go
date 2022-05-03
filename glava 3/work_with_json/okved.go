package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type OKVED struct {
	GlobalId int `json:"global_id"`
}

func main() {
	var (
		okveds []OKVED
		summa  int
		//buf    = new(bytes.Buffer)
	)
	file, err := os.Open("./glava 3/work_with_json/json_okved.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	enc := json.NewDecoder(file)
	enc.Decode(&okveds)
	for _, okved := range okveds {
		summa += okved.GlobalId
	}
	fmt.Println(summa)
}
