package main

import "fmt"

func rotateArray(data []int, k int) {
	n := len(data)
	reverseArray(data, 0, k-1)
	reverseArray(data, k, n-1)
	reverseArray(data, 0, n-1)
}

func reverseArray(data []int, start int, end int) {
	i := start
	j := end

	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

func main() {
	data := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(data)
	rotateArray(data, 2)
	fmt.Println(data)
}
