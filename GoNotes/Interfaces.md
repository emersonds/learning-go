Allows us to define a set of methods that different object types can implement. Similar to in Unity game dev, where I may have a `IDamageable<>` interface that any object can implement to call `TakeDamage()`.

Interfaces allow us to treat various types as the same type, allowing for more flexible/modular code.

For example, an interface to calculate the area of shapes:
```go
// Interface declaration
type Geometry interface {
	area() int  // area() is the method that can be called when a shape is given this interface
}

// area() implementation for a Circle
func (c Circle) area() int {
    return math.Pi * c.radius * c.radius
}

// area() implementation for a Square
func (s Square) area() int {
    return s.side * s.side
}

// Now we can access the area() method from the Geometry interface in a separate function, getArea()
func getArea(shape any) (int, error) {
	// _ is a lambda expression for a temporary shape type variable. Note the Geometry interface attached to shape, allowing us to call the area() method from the interface.
	// "ok" is a variable that determines if the shape has implemented the interface correctly - in other words if "_" is a valid shape.
    _, ok := shape.(Geometry); ok {
	  return shape.area(), nil // ok is nil here because the shape is valid
    } else {
      return 0, errors.New("unknown type") // shape does not implement the interface thus is an unknown type, and "ok" becomes an error message.
    }
}

```
[Go by Example: Interfaces](https://gobyexample.com/interfaces) seems to break these down better.

[Video Explaining Interfaces](https://www.youtube.com/watch?v=SX1gT5A9H-U) by Flo Woelki
```
(Comment from @minhleuc2346 on the above video)
In Go (Golang), there are no "classes" like in Python, Java, or C++.

Instead, Go uses:
- structs: These are used to group data together (like fields or attributes).
- methods: These are functions that are attached to a struct. They can work with the data inside the struct.

Go also has interfaces, which are just a list of method names. To "implement" an interface, a struct needs to have all the methods listed in that interface.

When you call a method from an interface, you can pass a variable of a struct type — as long as that struct has already implemented the needed methods.
```