Maps and Arrays hold a set amount of items, using `range` to work through each item one at a time:
```go
letters := []string{"A", "B", "C", "D"}
// Assign two new variables to hold the index and the value of each item in the array
// This is similar to for(int i=0) and arr[i]
for index, value := range letters {
  fmt.Println("Index:", index, "Value:", value)
}
```

Maps do not have an index, rather they use key-value pairs like a dictionary.
```go
addressBook := map[string]string{
  "John": "12 Main St",
  "Janet": "56 Pleasant St",
  "Jordan": "88 Liberty Ln",
}
// Key-Value pairs instead of array index values
for key, value := range addressBook {
  fmt.Println("Name:", key, "Address:", value)
}
```

Can use a blank identifier `_` with range to not have to print the index and value:
```go
menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

fmt.Println("The menu:")

// Prints each value without needing the index.
// Go throws an error if a variable is declared and not used
for _, value := range menu {
	fmt.Println(value);
}
```

Arrays in Go must be provided a number of elements, and that amount cannot change without declaring a new array. However, arrays do not need the elements to be assigned.
```go
var playerScores [4]int // An int array holding 4 ints with default 0 value
```

The Go compiler can automatically determine the length of an array using `...` in place of the number of elements if the value of the elements is already known.
```go
myItems := [...]string{"shoe", "watch", "toy"} // Compiler knows array has 3 elements

lotteryNumbers := [...]int{5, 48, 32, 1, 6}
fmt.Println(lotteryNumbers)    // Prints [5 48 32 1 6]
```

Arrays can have their elements changed at any index.
```go
myArray := [4]int{10, 24, 5, 47}
myArray[2] = 33
fmt.Println(myArray) // Prints [10 24 33 47]
```

### Slices
Similar to arrays but can change their size - similar to Lists in C#. The syntax for creating a slice is similar to creating an array, but the number of elements is not assigned:
```go
// Each of the following creates an empty slice
var numberSlice []int
stringSlice := []string{}

// The following creates a slice with elements
names := []string{"Kathryn", "Martin", "Sasha", "Steven"}
```

Slices can also be created from existing arrays. Slices made this way can contain the full array or specific elements of the array depending on how it is declared. The `:` operator in the number of elements determines the starting point (the left side of `:`) and the end point (right side of `:`) of the elements in the slice from the original array.
```go
// Regular int array declaration
array := [5]int{2, 5, 7, 1, 3}
// This is a slice of the whole array
sliceVersion := array[:]
fmt.Println(sliceVersion)
// [2 5 7 1 3]
// This is a slice containing the elements at indices 2, 3, and 4
partialSlice := array[2:5]
fmt.Println(partialSlice)
// [7 1 3]
// This is a slice containing the elements at indices 0 through 2
frontSlice := array[:3]
fmt.Println(frontSlice)
// [2 5 7]
// This is a slice containing elements at indices 2 through 4
backSlice := array[2:]
fmt.Println(backSlice)
// [7 1 3]
```

Slices can be modified the same way as arrays, using the index number.

`len()` is a function that returns the length of an array or slice, **how many elements it holds**:
```go
myArray := []int{1, 2, 4}
fmt.Println(len(myArray))    // 3
```

Slices are resizeable, so there is a difference between:
- Length: Current number of elements held
- **Capacity**: Number of elements the slice *can hold* before needing to resize itself
A slice's capacity can be found using the `cap()` function:
```go
slice := []string{"Fido", "Fifi", "FruFru"}
// The slice begins at length 3 and capacity 3
fmt.Println(slice, len(slice), cap(slice))
// [Fido Fifi FruFru] 3 3
slice = append(slice, "FroFro")
// After appending an element when the slice is at capacity
// The slice will double in capacity, but increase its length by 1
fmt.Println(slice, len(slice), cap(slice))
// [Fido Fifi FruFru FroFro] 4 6
```
It is important to note that appending an element to a slice when it is at capacity will *double the capacity* of the slice, but only increase its length by 1. This is not possible with arrays because additional values cannot be assigned when an array is full.

To add elements to a slice, we can use `append()`. This automatically handles the logic of adding to and resizing a slice.
```go
books := []string{"Tom Sawyer", "Of Mice and Men"}
fmt.Println(cap(books), len(books))    // 2 2
books = append(books, "Frankenstein")
books = append(books, "Dracula")
fmt.Println(cap(books), len(books))    // 4 4
fmt.Println(books)
// [Tom Sawyer Of Mice and Men Frankenstein Dracula]
```

Slices can be used to retain changes to an array when passing arrays to functions. Changes to the slice parameter will be permanent. Modifying a normal array parameter will not create permanent changes. Changes to the array parameter will only be local to the function.