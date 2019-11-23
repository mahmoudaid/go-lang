
GoLang Cheat Sheet
==================
## Table of Contents
- [Run example](#run-example)
- [Hello world](#hello-world)
- [Variables](#variables)
- [Constants](#constants)
- [Basic types](#basic-types)
	- [String](#string)
	- [Numbers](#numbers)
	- [Arrays](#arrays)
	- [Slices](#slices)
	- [Pointer](#pointer)
	- [Type conversions](#type-conversions)
- [Flow control](#flow-control)
	- [Conditional](#conditional)
	- [Switch](#switch)
	- [For-Range loop](#for-range-loop)
- [Functions](#functions)
	- [Lambdas](#lambdas)
	- [Multiple return types](#multiple-return-types)
	- [Named return values](#named-return-values)
- [Packages](#packages)
	- [Declaration](#declaration)
	- [Importing](#importing)
	- [Aliases](#aliases)
	- [Exporting names](#exporting-names)
- [Concurrency](#concurrency)
	- [Goroutines](#goroutines)
	- [Buffered channels](#buffered-channels)
	- [Closing channels](#closing-channels)
- [Error control](#error-control)
	- [Defer](#defer)
  - [Deferring functions](#deferring-functions)
- [Structs](#structs)
  - [Defining](#defining)
  - [Literals](#literals)
  - [Pointers to structs](#pointers-to-structs)
- [Methods](#methods)
  - [Receivers](#receivers)
  - [Mutation](#mutation)
- [References](#references)

# Run example
`$ go run cheatsheet.go`

# Hello world
```
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```

# Variables
The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level.
## Multiple variables declaration.
```
var c, python, java bool
```

## Variable declaration then assign value.
```
var var1 string;
var1 = "Hello var ex 1"
fmt.Println(var1)
```

## Variable declaration and assign value in one line.
```
var var2 = "Hello var ex 1"
fmt.Println(var2)
```

## Variable declaration and assign Shortcut (Infers type).
```
var3 := "Hello var ex 2"
fmt.Println(var3)
```

# Constants
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.
```
const Pi = 3.14
```

# Basic types
## String
```
str1 := "Hello"
```
Multiline string
```
str2 := `This is
Multiline string.`
```
## Numbers
### Typical types
```
num := 3 // integer
num := 3. // float64
num := 3 + 4i // complex128
num := byte('a')  // byte (alias for uint8)
```
## Arrays
Arrays have a fixed size.
```
// var numbers [5]int
numbers := [...]int{0, 0, 0, 0, 0}
```

## Slices
Slices have a dynamic size, unlike arrays.
```
slice := []int{2, 3, 4}
slice2 := []byte("Hello")
```

## Pointer
Pointers point to a memory location of a variable. Go is fully garbage-collected.
```
func main () {
  b := *getPointer()
  fmt.Println("Value is", b)
}
 
func getPointer () (myPointer *int) {
  a := 234
  return &a
}
```

## Type conversions
```
i := 2
f := float64(i)
u := uint(i)
```

# Flow control
## Conditional
Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
```
if day == "sunday" || day == "saturday" {
	fmt.Println("Weekend")
} else if day == "monday" && isTired {
	fmt.Println("Day off work")
} else {
	fmt.Println("Work")
}
```
A Tour of Go: [If](https://tour.golang.org/flowcontrol/5)

A condition in an if statement can be preceded with a statement before a ;.
```
if _, err := getResult(); err != nil {
  fmt.Println("Uh oh")
}
```
A Tour of Go: [If with a short statement](https://tour.golang.org/flowcontrol/6)

## Switch
```
switch day {
  case "sunday":
    fmt.Println("Visit family")
  case "saturday":
    fmt.Println("Shopping")

  default:
    fmt.Println("Work")
}
```
Go wiki: [Switch](https://github.com/golang/go/wiki/Switch)

## For-Range loop
```
names := []string{"mahmoud", "Zayed", "Jack","John","Jones"}
for i, val := range names {
	fmt.Printf("At position %d, the name %s is present\n", i, val)
}
```
Go by Example: [Range](https://gobyexample.com/range)

## For loop
Go has only one looping construct, the for loop.
```
for count := 0; count <= 10; count++ {
	fmt.Println("My counter is at", count)
}
```
A Tour of Go [For](https://tour.golang.org/flowcontrol/1)

# Functions
## Lambdas
Functions are first class objects.
```
myfunc := func() bool {
  return x > 10000
}
```

## Multiple return types
```
a, b := getMessage()
func getMessage() (a string, b string) {
  return "Hello", "World"
}
```

## Named return values
By defining the return value names in the signature, a return (no args) will return variables with those names.
```
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}
```
A Tour of Go: [Named return values](https://tour.golang.org/basics/7)

# Packages
## Declaration
Every package file has to start with package.
```
package hello
```

## Importing
```
import "fmt"
import "math/rand"
```
Or
```
import (
  "fmt"
  "math/rand"
)
```
Both are the same.

A tour of Go: [Importing](https://tour.golang.org/basics/1)

## Aliases
```
import r "math/rand"
r.Intn()
```

## Exporting names
Exported names begin with capital letters.
```
func Hello () {
  ···
}
```
A Tour of Go: [Exporting names](https://tour.golang.org/basics/3)

# Concurrency
## Goroutines
Channels are concurrency-safe communication objects, used in goroutines.
```
func main() {
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
}

func push(name string, ch chan string) {
  msg := "Hey, " + name
  ch <- msg
}
```
A Tour of Go: [Goroutines](https://tour.golang.org/concurrency/1), [Channels](https://tour.golang.org/concurrency/2)

## Buffered channels
Buffered channels limit the amount of messages it can keep.
```
ch := make(chan int, 2)
ch <- 1
ch <- 2
ch <- 3
// fatal error:
// all goroutines are asleep - deadlock!
```
A Tour of Go: [Buffered channels](https://tour.golang.org/concurrency/3)

## Closing channels
Closes a channel
```
ch <- 1
ch <- 2
ch <- 3
close(ch)
```
Iterates across a channel until its closed
```
for i := range ch {
  ···
}
```
Closed if ok == false
```
v, ok := <- ch
```
A Tour of Go: [Range and close](https://tour.golang.org/concurrency/4)

# Error control
## Defer
Defers running a function until the surrounding function returns. The arguments are evaluated immediately, but the function call is not ran until later.
```
func main() {
  defer fmt.Println("Done")
  fmt.Println("Working...")
}
```
Go blog: [Defer, panic and recover](https://blog.golang.org/defer-panic-and-recover)

## Deferring functions
Lambdas are better suited for defer blocks.
```
func main() {
  defer func() {
    fmt.Println("Done")
  }()
  fmt.Println("Working...")
}
```

# Structs
## Defining
```
type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X, v.Y)
}
```

## Literals
```
v := Vertex{X: 1, Y: 2}

// Field names can be omitted
v := Vertex{1, 2}

// Y is implicit
v := Vertex{X: 1}
```

## Pointers to structs
Doing v.X is the same as doing (*v).X, when v is a pointer.
```
v := &Vertex{1, 2}
v.X = 2
```
# Methods
## Receivers
There are no classes, but you can define functions with receivers.
```
type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
 
v: = Vertex{1, 2}
v.Abs()
```
A Tour of Go: [Methods](https://tour.golang.org/methods/1)

## Mutation
By defining your receiver as a pointer (*Vertex), you can do mutations.
```
func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}
 
v := Vertex{6, 12}
v.Scale(0.5)
// `v` is updated
```
A Tour of Go: [Pointer receivers](https://tour.golang.org/methods/4)

# References
- [A tour of Go](https://tour.golang.org/)
- [Golang wiki](https://github.com/golang/go/wiki)
- [Awesome Go](https://awesome-go.com/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [JustForFunc Youtube](https://www.youtube.com/channel/UC_BzFbxG2za3bp5NRRRXJSw)
- [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)
- [cheat-sheets.org](http://www.cheat-sheets.org/saved-copy/go-lang-cheat-sheet-master.20181212/golang_refcard.pdf)
- [devhints.io](https://devhints.io/go)

