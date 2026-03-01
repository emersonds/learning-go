Useful for grouping related data together and writing functions that work with structs. Struct instances can be declared by assigning all values of a struct, some values of a struct, or omitting all values of a struct to create an instance using default values.
```go
// Struct declaration
type StructName struct {
    type1 int
    type2 string
    type3 float32
}

// Associated struct function that accesses the struct (note the struct receiver)
func (s StructName) StructFunction() int {
	return s.type1
}

// Associated function that uses a pointer to directly access and modify function values
func (s *StructName) StructFunction2(newInt int) {
	s.type1 = newInt
}

func main() {
    // Create a new instance of struct StructName
	structInstance := StructName{type1: 1, type2: "str", type3: "0.32"}
    
    // If fields are omitted, Go will use default values
    instance2 := StructName{type1: 5} // {5, "", 0.0}
    
    // All fields can be omitted
    instance3 := StructName{} // {0, "", 0.0}
    
    // Because values are assigned from left to right according to how the fields are defined in the struct from top to bottom, we also don't have to label our fields in the instance declaration
    instance4 := StructName{10, "str2", 0.1} // Same as StructName{type1: 10, type2: "str2", type3: 0.1}
}
```

Individual fields of a struct instance can be accessed using the format `instance.field`:
```go
john := Student{firstName: "John", lastName: "Smith", age: 15}
fmt.Println(john.age) // Prints 15

john.age = 16  // "age" field in Student struct instance john changed to 16
```

