package main

import "fmt"

// Vertex Struct Defining.
type Vertex struct {
	X int
	Y int
}

func main() {
	docSection("Hello world")
	// Hello world
	fmt.Println("Hello world")

	docSection("Variables")
	// Variables
	//// Variable declaration then assign.
	var var1 string
	var1 = "Hello var ex 1"
	fmt.Println(var1)

	//// Variable declaration and assign in one line.
	var var2 = "Hello var ex 1"
	fmt.Println(var2)

	//// Variable declaration and assign Shortcut (Infers type).
	var3 := "Hello var ex 2"
	fmt.Println(var3)

	docSection("Constants")
	// Constants
	const Pi = 3.14
	fmt.Println(Pi)

	docSection("Basic types")
	// Basic types
	docSubSection("String")
	// String
	str1 := "Hello"
	fmt.Println(str1)

	str2 := `This is
	Multiline string.`
	fmt.Println(str2)

	docSubSection("Numbers")
	num1 := 3 // integer
	fmt.Println(num1)

	num2 := 3. // float64
	fmt.Println(num2)

	num3 := 3 + 4i // complex128
	fmt.Println(num3)

	docSubSection("Arrays")
	// var numbers [5]int
	numbers := [...]int{0, 0, 0, 0, 0}
	fmt.Println(numbers)

	docSubSection("Slices")
	// Slices have a dynamic size, unlike arrays.
	slice := []int{2, 3, 4}
	fmt.Println(slice)
	slice2 := []byte("Hello")
	fmt.Println(slice2)

	docSubSection("Pointers")
	// Pointers
	b := *getPointer()
	fmt.Println("Value is", b)

	docSubSection("Type conversions")
	// Type conversions
	i := 2
	f := float64(i)
	fmt.Println(f)
	u := uint(i)
	fmt.Println(u)

	docSection("Flow control")
	// Flow control
	docSubSection("Conditional")
	// Conditional
	day := "sunday"
	isTired := true
	if day == "sunday" || day == "saturday" {
		fmt.Println("Weekend")
	} else if day == "monday" && isTired {
		fmt.Println("Day off work")
	} else {
		fmt.Println("Work")
	}

	docSubSection("Switch")
	// Switch
	switch day {
	case "sunday":
		fmt.Println("Visit family")
	case "saturday":
		fmt.Println("Shopping")

	default:
		fmt.Println("Work")
	}

	docSubSection("For-Range loop")
	// For-Range loop
	names := []string{"mahmoud", "Zayed", "Jack", "John", "Jones"}
	for i, val := range names {
		fmt.Printf("At position %d, the name %s is present\n", i, val)
	}

	docSubSection("For loop")
	// For loop
	for count := 0; count <= 10; count++ {
		fmt.Println("My counter is at", count)
	}

	docSection("Functions")
	// Functions
	docSubSection("Lambdas")
	// Lambdas
	x := 100
	myfunc := func() bool {
		return x > 10000
	}
	fmt.Printf("%t\n", myfunc())

	docSubSection("Multiple return types")
	// Multiple return types
	m, n := getMessage()
	fmt.Println(m, n)

	docSubSection("Named return values")
	// Named return values
	fmt.Println(split(10))

	docSection("Concurrency")
	// Concurrency
	docSubSection("Goroutines")
	// Goroutines

	// A "channel"
	ch := make(chan string)

	// Start concurrent routines
	go push("Moe", ch)
	go push("Larry", ch)
	go push("Curly", ch)

	// Read 3 results
	// (Since our goroutines are concurrent,
	// the order isn't guaranteed!)
	fmt.Println(<-ch, <-ch, <-ch)

	docSection("Error control")
	// Error control
	docSubSection("Defer")
	// Defer
	defer fmt.Println("Defer Done")
	fmt.Println("Working...")

	docSubSection("Deferring functions")
	// Deferring functions
	defer func() {
		fmt.Println("Defer Done")
	}()
	fmt.Println("Working...")

	docSection("Struct")
	// Struct
	docSubSection("defining")
	// defining
	v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X, v.Y)
}

/**
 * Pointer.
 */
func getPointer() (myPointer *int) {
	a := 234
	return &a
}

/**
 * Helper function for styling the output.
 */
func docSection(name string) {
	sp := "********************************************"
	fmt.Println(sp)
	fmt.Println("\t\t" + name)
	fmt.Println(sp)
}

/**
 * Helper function for styling the output.
 */
func docSubSection(name string) {
	sp := "----------------------"
	fmt.Println(sp)
	fmt.Println(name)
	fmt.Println(sp)
}

/**
 * Multiple return types
 */
func getMessage() (a string, b string) {
	return "Hello", "World"
}

/**
 * Named return values
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/**
 * Concurrency Goroutines.
 */
func push(name string, ch chan string) {
	msg := "Hey, " + name
	ch <- msg
}
