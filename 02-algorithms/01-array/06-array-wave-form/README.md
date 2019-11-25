# Problem:
Given an array, arrange its elements in a wave form such that odd elements are lesser than its neighboring even elements.
# Solution:
## First Solution:
Compare even index values with it's neightbor odd index values and swap if odd index is not smaller than even index.
```
package main

import "fmt"

func swap(data []int, i int, j int) {
	data[i], data[j] = data[j], data[i]
}

func waveArray(data []int) {
	size := len(data)

	for i := 1; i < size; i += 2 {
		if (i-1) > 0 && data[i] > data[i-1] {
			swap(data, i, i-1)
		}

		if (i+1) < size && data[i] > data[i+1] {
			swap(data, i, i+1)
		}
	}
}

func main() {
	data := []int{8, 1, 2, 3, 4, 5, 6, 4, 2}
	fmt.Println(data)

	// Solution 1
	waveArray(data)
	fmt.Println(data)
}
```
## Second solution:
Sort the array then swap i and i + 1 index value so the array will form the wave.
```
package main

import (
	"fmt"
	"sort"
)

func swap(data []int, i int, j int) {
	data[i], data[j] = data[j], data[i]
}

func waveArray2(data []int) {
	size := len(data)
	sort.Ints(data)
	for i := 0; i < size-1; i += 2 {
		swap(data, i, i+1)
	}
}

func main() {
	data := []int{8, 1, 2, 3, 4, 5, 6, 4, 2}
	fmt.Println(data)

	// Solution 2
	waveArray2(data)
	fmt.Println(data)
}
```