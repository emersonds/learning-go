Go is a *pass-by-value* language. *Passing* functions the *value* of an argument. Go compiler strictly uses the **value** of the argument rather than the argument itself, thus changes that take place in the function stay within the function.

With *addresses*, *pointers*, and *dereferencing*, we can  change values from different scopes.

### Addresses
Using `&` we can find a variable's address, which is formatted in hexadecimal:
```go
x := "My very first address"
fmt.Println(&x) // Prints 0x414020
```

### Pointers
Pointers are variables used to *store* addresses. They also have a data type for the addresses' value.
```go
var pointerForInt *int // Normal/default variable declaration

pointerForInt := &x // Short declaration
```
The above variable declaration tells us that `pointerForInt` will store the address of a variable of `int` data type, identified by the `*`. 
For example:
```go
// Declare a pointer for an int type variable
var pointerForInt *int

minutes := 525600  // int type

// Assign pointerForInt to the address of minutes
pointerForInt = &minutes

// Print the address
fmt.Println(pointerForInt) // Prints 0xc000018038
```

### Dereferencing
Using a pointer, we can access a variable's address and change its value. This is called *dereferencing* or *indirecting*. Using the `*` operator on a pointer allows us to assign a new value to it, like so:
```go
x := 7 // Declare int x with value of 7
pointerForX := &x // Declare x pointer

*pointerForX = 10 // Assign a new value to x by dereferencing

fmt.Println(x) // Prints 10
```

Knowing how to change the value of an address with a pointer, we can now change the value of variables in different scopes! We do this by passing the value of a pointer - which is an address - and dereferencing it to change its value:
```go
// Adds 100 to a pointer passed as an argument
// This function is called with the address of the variable we are adding 100 to, creating a pointer with the address. addHundred(&x)
func addHundred (numPtr *int) {
  // Dereference the pointer to change its value
  *numPtr += 100
}

func main() {
  x := 1
  // Pass the address of x to addHundred() for the pointer parameter
  addHundred(&x)
  fmt.Println(x) // Prints 101
}
```

In summary, if we want to change the value of a variable outside of its current scope, we have to 1) pass the ***address*** of the variable to a function with a *pointer* parameter, which creates a pointer with the passed address. 2) We ***dereference*** the pointer to change the value of the address, ultimately changing the original variable outside of its scope.

Bonus: A pointer that has not been assigned a value has a default value of `<nil>`
```go
var pointerToNothing *string // or *int, *float, *bool, etc

fmt.Println("pointerToNothing:", pointerToNothing) // Prints "pointerToNothing: <nil>"
```
