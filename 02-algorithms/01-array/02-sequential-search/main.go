package main

import "fmt"

func sequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(sequentialSearch(data, 7))
	fmt.Println(sequentialSearch(data, 8))
}
