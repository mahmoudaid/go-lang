package main

import "fmt"

func indexArray(data []int) {
	size := len(data)

	for i := 0; i < size; i++ {
		curr := i
		value := -1

		for data[curr] != -1 && data[curr] != curr {
			temp := data[curr]
			data[curr] = value
			value = temp
			curr = value
		}

		if value != -1 {
			data[curr] = value
		}

	}
}

func indexArray2(data []int) {
	size := len(data)

	for i := 0; i < size; i++ {
		for data[i] != -1 && data[i] != i {
			temp := data[i]
			data[i] = data[temp]
			data[temp] = temp
		}
	}
}

func main() {
	data := []int{8, -1, 6, 1, 9, 3, 2, 7, 4, -1}
	fmt.Println(data)

	// Solution 1
	indexArray(data)
	fmt.Println(data)

	// Solution 2
	indexArray2(data)
	fmt.Println(data)
}
