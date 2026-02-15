*Learning with [Codecademy](https://www.codecademy.com/courses/learn-go/)*

`.go` files can be organized into packages and imported to another file. For example, writing a calculator program, files for calculation could be stored in `package calc` and files for input/output stored in `package io`.

A **package declaration** tells the Go compiler which package this file belongs to, where `package main` will be compiled into an executable on `go build main`.

An **import statement** allows the use of code from other packages: `import "fmt"`. Go's compiler will throw an error if a package is imported but unused.

`func` declares a function:
```go
func main () {
	fmt.Println("Hello World")
}
```

`.go` files can be built and run using the commands `go build main.go` and `./main`, where `main` is the resulting executable file created by `go build`. A `.go` file can also be run using `go run main.go` but the `go run` command **does not build an executable file**.
### Go Doc
`go doc` is used to extract and view documentation from `.go` files. View package docs by running `go doc <package name>` or info about a package function by running `go doc <package name>.<function>`

`go doc fmt` or `go doc fmt.Println`

[Go's official documentation](https://go.dev/doc/)

### Example function to print an ASCII dog
```go
package main

import f "fmt"

func main () {
	f.Printf(" __ _\n" +

		"o'')}____//\n" +

		" `_/ )\n" +

		" (_(_/-(_/\n")

}
```
```
  __      _
o'')}____//
 `_/      )
 (_(_/-(_/
```

### Variables
```go
var lengthOfSong uint16
var isMusicOver bool
var songRating float32 = 0.52c
```

Go will have a compiler error if a variable is defined but not used. File\:line\/column
`./main.go:4:7: lengthOfSong declared and not used`

**Short Declaration Operator** `:=` allows variable declaration without explicitly stating its type. This infers the type of a variable.
```go
daysOnVacation := 10
```
Floats created this way are `float64` while integers are either `int32` or `int64` depending on if the computer architecture is running on 32-bit or 64-bit.

Go supports multiple variable declaration
```go
var part1, part2 string
part1 = "To be..."
part2 = "Not to be..."

// Using short declaration
quote, fact := "Bears, Beets, Battlestar Galactica", true
```

`string` types should be assigned with double quotes (`""`). Single quotes (`''`) are used exclusively for defining a single character (also called **a rune**): `'a'`.

### Conditionals
Similar to C#, `if-else` conditionals use curly brackets. `else` has to be preceded by the closing curly bracket from the `if` statement:
```go
isTrue := true

if !isTrue {
	fmt.Print("False")
} else {
	fmt.Print("True")
}
```

`switch` blocks are useful for checking multiple values at once.
```go
switch name {
	case "Joe":
		fmt.Println("Mama")
	case "Billy":
		fmt.Println("Bob")
	default:
		fmt.Println("Nice name")
}
```

`if` and `switch` statements support **scoped short declaration statements (`:=`)**. This means we can declare a variable in our conditional statements and check against them.
```go
// If statement short declaration
x := 8
y := 9
if product := x * y; product > 60 {
  fmt.Println(product, "  is greater than 60")
}

// Switch statement short declaration
switch season := "summer" ; season {
case "summer":
  fmt.Println("Go out and enjoy the sun!")
}
```
However, these variables are only scoped to the `if` or `switch` statement they are declared in.
```go
x := 8
y := 9
if product := x * y; product > 60 {
	fmt.Println(product, "is greater than 60")
}

 // "Undefined" error, product does not exist in this scope!!
product *= 0.0887
```

### Random Numbers and Seeding
Go has a `"math/rand"` library for generating random numbers. Random integers can be generated using `rand.Intn(n)` where a random number is chosen from 0 to `n-1` (Upper bound exclusive).

However, random numbers must be seeded to prevent the same number from being selected every time the function is called.

A common way to do this is by incorporating the `"time"` package, because the current time will always be different, thus always giving us a unique number:
```go
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())  // Difference in time since January 1st, 1970
  fmt.Println(rand.Intn(100))
}
```

### Functions
Works the same as C-like and other programming languages. Return type follows the function name rather than preceding it.
```go
func getLengthOfCentralPark() int32 {
  var lengthInBlocks int32
  lengthInBlocks = 51
  return lengthInBlocks
}

func main () {
	length := getLengthOfCentralPark()
}
```

Just like declaring variables, `var x int32`, function parameters have the type after the variable declaration. With multiple parameters of the same type, the type declaration can be at the end of the parameter list.
```go
// This function takes two int32 parameters and returns an int32
func multiplier(x, y int32) int32 {
  return x * y
}
```

Functions can return multiple values. See the following example where `func GPA()` returns a string and a float32:
```go
// Calculates GPA and returns two values, the letter grade (string) and the average grade (float32)
func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
  averageGrade := (midtermGrade + finalGrade) / 2
  var gradeLetter string
  if averageGrade > 90 {
    gradeLetter = "A"
  } else if averageGrade > 80 {
    gradeLetter = "B"
  } else if averageGrade > 70 {
    gradeLetter = "C"
  } else if averageGrade > 60 {
    gradeLetter = "D"
  } else {
    gradeLetter = "F"
  }

  return gradeLetter, averageGrade 
}

func main() {
  var myMidterm, myFinal float32
  myMidterm = 89.4
  myFinal = 74.9
  var myAverage float32
  var myGrade string
  // Note how GPA() returns two values
  // Use two variables to assign each value
  myGrade, myAverage = GPA(myMidterm, myFinal)
  fmt.Println(myAverage, myGrade) // Prints 82.12 B
}
```

`defer` keyword can be used to delay a function call to the end of the current scope. `defer` tells Go to run a function at the end of the current function. This is useful for logging, file writing, and other utilities.

In this example, `defer` is called right at the beginning of the function, but does not print anything until the function ends (in this case, a value is returned).
```go
func calculateTaxes(revenue, deductions, credits float64) float64 {
  defer fmt.Println("Taxes Calculated!")
  taxRate := .06143
  fmt.Println("Calculating Taxes")

  if deductions == 0 || credits == 0 {
    return revenue * taxRate
  }
  

  taxValue := (revenue - (deductions * credits)) * taxRate
  if taxValue >= 0 {
    return taxValue
  } else {
    return 0
  }
}
```