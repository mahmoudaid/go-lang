# Problem:
Write a method that will return the sum of all elements of the integer Array, given Array as an input argument.

# Solution:
```
import "fmt"

func sumArray(data []int) int {
	size := len(data)
	total := 0

	for i := 0; i < size; i++ {
		total += data[i]
	}

	return total
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("The sum of all values in array:", sumArray(data))
}
```
