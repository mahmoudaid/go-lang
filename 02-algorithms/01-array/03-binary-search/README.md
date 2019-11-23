# Problem:
If we want to search some value in sorted array, a binary search can be used, Write a method which will search sorted array for some given value.

# Solution:
```
package main

import "fmt"

func binarySearch(data []int, value int) bool {
	low := 0
	var mid int
	hight := len(data) - 1

	for low <= hight {
		mid = (hight + low) / 2
		if value == data[mid] {
			return true
		} else {
			if data[mid] < value {
				low = mid + 1
			} else {
				hight = mid - 1
			}
		}
	}

	return false
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(binarySearch(data, 7))
	fmt.Println(binarySearch(data, 8))
}
```