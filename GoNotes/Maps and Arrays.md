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