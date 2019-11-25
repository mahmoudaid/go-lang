# Problem:
Given array of size N , containing elements from 0 to N-1. All values 0 to N-1 are present in array and  if they are not there than -1 is there to take its place. Arrange values of array so that value i is stored at data[i]
# Solution:
## First solution:
For each value pick the element and then put it into its proper place and the element which is in its proper position then pick it and repeat the process again.

```
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

func main() {
	data := []int{8, -1, 6, 1, 9, 3, 2, 7, 4, -1}
	fmt.Println(data)

	// Solution 1
	indexArray(data)
	fmt.Println(data)
}

```

## Second solution:
For each index, we will pick the value at an index and put that value at its proper position.
```
package main

import "fmt"

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

	// Solution 2
	indexArray2(data)
	fmt.Println(data)
}

```