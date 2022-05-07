package main

import "fmt"

func main() {
	goroutinesWork()
	//inputStream := make(chan string)
	//outputStream := make(chan string)
	//go removeDuplicates(inputStream, outputStream)
	//
	//go func() {
	//	defer close(inputStream)
	//
	//	for _, r := range "112334456" {
	//		inputStream <- string(r)
	//	}
	//}()
	//
	//for x := range outputStream {
	//	fmt.Print(x)
	//}
}

func task(channel chan int, N int) {

	i := N + 1
	channel <- i
}

func task2(channel chan string, N string) {
	for i := 0; i < 5; i++ {
		channel <- N + " "
	}
}

func removeDuplicates(inputStream, outputStream chan string) {
	defer close(outputStream)
	c := make([]string, 0)
	for i := range inputStream {
		if !inSlice(i, c) {
			c = append(c, i)

			outputStream <- i
		}
	}
}

func inSlice(i string, c []string) bool {
	if len(c) > 0 {
		for _, el := range c {
			if i == el {
				return true
			}
		}
		return false
	} else {
		return false
	}
}

func goroutinesWork() {
	channel := make(chan int)
	go func(channel chan int) {
		work()
		close(channel)
	}(channel)
	<-channel
}

func work() {
	fmt.Println("hsgkjlafd")
}
