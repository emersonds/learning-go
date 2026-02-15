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
	// No clue was "ok" is yet
    _, ok := shape.(Geometry); ok {
      return shape.area(), nil
    } else {
      return 0, errors.New("unknown type")
    }
}

```
[Go by Example: Interfaces](https://gobyexample.com/interfaces) seems to break these down better.