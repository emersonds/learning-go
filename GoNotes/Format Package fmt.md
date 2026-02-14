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
| Verb     | Usage                                       | Example                                         |
| -------- | ------------------------------------------- | ----------------------------------------------- |
| %v       | The value in default format                 | `("%v", num)` prints value of num, `1`          |
| %T       | The type of the value                       | `("%T", num)` prints type of num, `int`         |
| %t       | The word true or false                      | `("%t", isNum)` prints `true`                   |
| %d       | Interpolates `int` to `string` in base10    | `("%d", num)` prints `1`                        |
| %f or %F | Decimal point, as well as decimal precision | `("%.2f", dec)` prints `3.80` where `dec = 3.8` |
| %s       |                                             |                                                 |
| %q       |                                             |                                                 |
