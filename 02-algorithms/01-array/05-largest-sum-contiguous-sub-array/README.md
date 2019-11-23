# Problem:
Given an Array of positive and negative integers, find a contiguous subarray whose sum (sum of elements) is maximum.
# Solution:
```
package main

import "fmt"

func maxSubArray(data []int) int {
	size := len(data)
	maxSoFar := 0
	maxEndingHere := 0

	for i := 0; i < size; i++ {
		maxEndingHere += data[i]
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}

		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}

	return maxSoFar
}

func main() {
	data := []int{1, -2, 3, 4, -4, 6, -14, 6, 2}
	fmt.Println("Max sub array sum", maxSubArray(data))
}
```