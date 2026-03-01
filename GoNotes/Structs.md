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

If we want to use many instances of a struct, we can use an array of structs. Contents of a struct array can be accessed and modified, as well as the instance fields for each index.
```go
type Point struct {
	x int
	y int
}

// A slice of Point structs
points := []Point{{1, 1}, {7, 27}, {9, 25}}

// Can also declare an array with existing instances
a = {1, 1}
b = {7, 27}
c = {12, 7}
d = {9, 25}

points := []Point{a, b, c, d}

fmt.Println(points[0]) // Prints {1, 1}

points[0].x = 2
points[0].y = 3
fmt.Println(points[0].x) // Prints 2
```

Structs can also be nested to make complex groups of fields easier to work with.
```go
type Name struct {
	firstName string
	lastName string
}

type Employee struct {
	name Name  // Nested Name struct will contain firstName and lastName
	age int
	title string
}

carl := Employee{Name{"Carl", "Carlson"}, 32, "Engineer"}

// Nested structs can be accessed by chaining together the field accesses
// carl(instance).name(nested struct Name).lastName(field from nested Name struct)
fmt.Println(carl.name.lastName)  // Prints "Carlson"
```

When nesting structs, a nested struct can also be defined "anonymously".  This allows us to access a field from the name struct without having to chain field accesses. The only downside of this is that we cannot have two structs declared anonymously, because we won't know which field is being accessed. Anonymous fields lead to cleaner code because it is easier to read.
```go
type Name struct {...} // see above
type Employee struct {
	Name
	age int
	title string
}

carl := Employee{Name{"Carl", "Carlson"}, 32, "Engineer"}
fmt.Println(carl.firstName)  // Prints "Carl"
```

