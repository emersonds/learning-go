[`fmt` docs](https://pkg.go.dev/fmt)

`Println()`, `Print()`, `Printf()` used to format and print data.

`Sprint()`, `Sprintln()`, `Sprintf()` used to format data but does not print to terminal.

`Scan()` used to get user input.

### `Println()`
`fmt.Println()` prints its arguments with an included space in between each argument and adds a line break at the end.
```go
fmt.Println("Println", "formats", "really well")
fmt.Println("Right?")
```
```
Println formats really well
Right?
```

### `Print()`
`fmt.Print()` does not have the default formatting `fmt.Println()` has, and does not include spacing between arguments or a line break after printing.
```go
fmt.Print("The answer is", ": ")
fmt.Print("12")
```
```
The answer is: 12
```

Both `Print()` and `Println()` are both useful for concatenating strings.

### `Printf()`
`fmt.Printf()` is useful for interpolating strings: leave placeholders and use values to fill placeholders. Use `go doc fmt` to see placeholders (called *verbs* in Go, located at the top of `go doc` output).  For example, `%v` gets a value.
```go
guess := "C"
fmt.Printf("Is %v your final answer?", guess)
// Prints: Is C your final answer?
```

#### Examples of Verbs
| Verb     | Usage                                              | Example                                         |
| -------- | -------------------------------------------------- | ----------------------------------------------- |
| %v       | The value in default format                        | `("%v", num)` prints value of num, `1`          |
| %T       | The type of the value                              | `("%T", num)` prints type of num, `int`         |
| %t       | The word true or false                             | `("%t", isNum)` prints `true`                   |
| %d       | Interpolates `int` to `string` in base10           | `("%d", num)` prints `1`                        |
| %f or %F | Decimal point, as well as decimal precision        | `("%.2f", dec)` prints `3.80` where `dec = 3.8` |
| %s       | Uninterpreted bytes of string or slice             |                                                 |
| %q       | Double-quoted string safely escaped with Go syntax |                                                 |
#### `Sprint()` and `Sprintln()`
Similar to `Print()` and `Println()`, but these do not print to the terminal. Instead, they can be used to format strings, like variables:
```go
grade := "100"
compliment := "Great job!"
teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

fmt.Print(teacherSays)
// Prints: You scored a 100 on the test! Great job!
```
`Sprint()` behaves like `Print()`, where no spaces and line breaks are added, while `Sprintln()` behaves like `Println()` and adds the default formatting.

#### `Sprintf()`
Behaves like `Printf()`, allowing the use of verbs. `Sprintf()` returns its value instead of printing it.
```go
correctAns := "A"
answer := fmt.Sprintf("And the correct answer is… %v!", correctAns)

fmt.Print(answer) // Prints: And the correct answer is… A!
```

#### `Scan()`
Used for getting user input. Be mindful of the `&` (address) before the variable in the function.
```go
fmt.Println("How are you doing?") 

var response string 
fmt.Scan(&response)

fmt.Printf("I'm %v.", response) 
```
*Note: `Scan()` does not allow whitespace, such as spaces, in the arguments. It will cut off everything after the first word if there is a space.*

