`for` loops have similar syntax to C# and other familiar languages
```go
for number := 0; number < 5; number++ {
  fmt.Print(number, " ")
}
// Output: 0 1 2 3 4
```

`for` can be used for definite and indefinite loops in Go. Rather than having a `while` keyword, we can simply change the condition
```go
// Loop while number is less than 5
for number < 5 {
	fmt.Println(number)
	number++
}

// Loop while number is not 56
for number != 56 {
	number = getNum()
}

// Indefinite loop with no end
for {
	// Loop body logic runs forever
}
// This code is never reached
```

`break` can be used to stop the loop at a specific iteration, and `continue` can be used to skip to the next iteration.
```go
// Example of break
// Stops loop when a dog is found in array
animals := []string{"Cat", "Dog", "Fish", "Turtle"}
for index := 0; index < len(animals); index++ {
  if animals[index] == "Dog" {
    fmt.Println("Found the perfect animal!")
    break // Stop searching the array
  }
}

// Example of continue
// Skips to next iteration, so "you ate the <color> jellybean" does not print
jellybeans := []string{"green", "blue", "yellow", "red", "green", "yellow", "red"}
for index := 0; index < len(jellybeans); index++ {
  if jellybeans[index] == "green" {
    continue
  }
  fmt.Println("You ate the", jellybeans[index], "jellybean!")
}
```