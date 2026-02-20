Allows us to define a set of methods that different object types can implement. Similar to in Unity game dev, where I may have a `IDamageable<>` interface that any object can implement to call `TakeDamage()`. A type *must implement all signatures present in the interface*.

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

Interfaces are implemented implicitly, meaning a type does not have to declare that it is using the interface. If a type has methods with the same signatures as the interface, it is considered an implementation of the interface.
```go
type Shape interface {
	GetArea() float64
	GetPerimeter() int
}

// Square implements the interface
type Square struct {
	GetArea() float64
	GetPerimeter() int
}

// Does not implement the interface
// Struct is missing  GetArea() float64
type Triangle struct {
	GetPerimeter() int
}
```

### Receivers
Methods are bound to *Receiver Types*, which can be value or pointer types, which affects how methods interact with data.
- **Value receivers** operate on *copies* of the data, meaning the original data is not modified.
```go
// Here, the receiver is `ms` of type `MaintenanceService`.
func (ms MaintenanceService) Service(vehicleID int) string { 
    location := ms.location
    return fmt.Sprintf("Vehicle %d has been serviced at %s.", vehicleID, location)
    
    // This will only increment ms.serviceCount in this function block
	ms.serviceCount++
}
```
- **Pointer Receivers** modify the original data because they reference the  *memory address*.
``` go
// Here, the receiver is a pointer receiver `ms` of type `MaintenanceService`.
func (ms *MaintenanceService) Service(vehicleID int) string { 
    location := ms.location
    return fmt.Sprintf("Vehicle %d has been serviced at %s.", vehicleID, location)
    
    // The original ms data will be incremented now because the pointer reference accesses its memory address.
	ms.serviceCount++ 
}
```

Receivers appear in their own argument list between `func` and the method name, and should typically be named with only one or two letters:
```go
// Func GetArea takes a value receiver, s, of type Square; has no parameters; and returns a  float64
func (s Square) GetArea() float64 {
	return s.x * s.y
}
```

In summary, interfaces are declared implicitly. As long as a type implements all methods specified in the interface, Go considers the interface implemented within the type. Receivers determine interface implementation separately based on their type: value receivers or pointer receivers.

### Utilizing interfaces
Another way to use interfaces is by declaring a variable of the interface type and assigning it to concrete types.
```go
type Geometry interface {
	GetArea()
}

type Square struct {
	length int
}

func (s Square) GetArea() int {
	return length * length
}

type Rectangle struct {
	length int
	width int
}

func (r Rectangle) GetArea() int {
	return length * width
}

func main () {
	var shape Geometry
	
	// Square and Rectangle implement Geometry interface
	shape = Square{ length: 4 }
	shape = Rectangle{ length: 3, width: 2 }
	
	// Because Square and Rectangle implement Geometry interface, we can call GetArea() from the interface
	// This will call the GetArea() method of whichever shape it currently represents (square or rectangle)
	area := shape.GetArea()
	
	fmt.Println("Area: ", area)  // Prints 6
}
```

### Best Practice: Programming to an Interface
It is better to program to an interface rather than a specific implementation. Using generic interfaces instead of concrete types keeps code less dependent on specific implementation, reduces code complexity and duplication, and enables easier updates and growth without significant refactoring.

### Type Assertions
Useful for getting or verifying the concrete type of an interface variable. Type assertions recover the concrete type of a variable with syntax `x.(T)`, with `x` being the interface and `T` being the type we are asserting.
- Safe Type Assertion: `value, ok := x.(T)` returns a Boolean `ok`  that checks if the assertion was successful, avoiding runtime panics.
- Unsafe Type Assertion: `value := x.(T)` attempts to retrieve the type and will cause panics at runtime if the interface `x` is not of type `T`.
	- Unsafe type assertions exist for scenarios where we are certain of what the type will be and want to avoid the overhead of checking the type and verifying with the `ok` bool.
Note: a runtime panic in Go is a critical error that halts normal program execution.
```go
var mediaPlayer Media = Audio{ duration: 120 }

// Safe type assetion
if media, ok := mediaPlayer.(Audio); ok {
	fmt.Println("This is media with duration:", media.duration)
} else {
	fmt.Println("Fail!")
}
```

### Type Switches
Extension of type assertions that allows us to handle multiple types safely in a single action. For example, if our interface is representing more than one concrete type:
```go
// "vehicle" is a variable assigned the Vehicle interface type
// v is assigned the result of the type assertion on vehicle
// Switch statement checks v against each case for each type implementation of Vehicle interface, in this case Car or Bike
switch v := vehicle.(type) {

case Car:
	fmt.Println("Car model: ", v.model)
case Bike:
	fmt.Println("Bike color: ", v.color)
default:
	fmt.Println("Unknown vehicle type")
}
```

### Empty Interfaces
`interface{}` can hold values of any type, because all types in Go implicitly implement at least zero methods, satisfying the conditions of the empty interface. The `any` keyword can be used in place of `interface{}`.

Empty interfaces are especially useful for functions that accept parameters of various types. For example, a function that processes different types of data and performs type-specific operations:

```go
var value interface{}

func processData(data any) {
	switch v := data.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Integer")
	case bool:
		fmt.Println("Boolean")
	default:
		fmt.Println("Unexpected type")
	}
}
```

Using empty interfaces comes with the risk of losing compile-time type checking, making it harder to ensure type safety. Use type assertions or switches to safely retrieve the underlying type of a value stored in an empty interface.

### Best Practices: Keep Interfaces Small
Interfaces should be small to keep them simple (KISS), flexible to allow mixing and matching of features, and to be combined into larger interfaces for more complex interfaces.
```go
// Creating bigger interfaces with smaller interfaces
type Reader interface {
    Read()
}

type Writer interface {
    Write()
}

// A type that implements ReadWriter must satisfy both Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
```

Combining interfaces allows for a more modular approach where each type only needs to implement the methods that are relevant to its role. Also allows for keeping parts of a system independent, only interacting with the necessary interfaces required to do their jobs.

### Best Practices: Behavior Over Implementation
Interfaces are best used to describe what a program should do rather than *how* a program should do it.
```go
type Notifier interface {
	Notify(message string)
}
```
The `Notifier` interface in this case allows for a flexible interface for notification methods. It tells the program *what* to do, notify something, but now *how* to do it. This interface can be implemented by multiple notification methods, such as SMS, email, or push notifications.
```go
package main

import "fmt"

func main() {

	var notifier Notifier
	
	notifier = EmailNotifier{}
	notifier.Notify("Hello via Email!")
	
	notifier = SMSNotifier{}
	notifier.Notify("Hello via SMS!")
}

// Define behavior-focused and abstract interface
type Notifier interface {
	Notify(message string)
}

  

// Implement the interface in different ways
type EmailNotifier struct{}

// EmailNotifier type implementation of Notifier interface
func (e EmailNotifier) Notify(message string) {
	fmt.Println("Sending email with message:", message)
}

type SMSNotifier struct{}

// SMSNotifier type implementation of Notifier interface
func (s SMSNotifier) Notify(message string) {
	fmt.Println("Sending SMS with message:", message)
}
	case Car:
		fmt.Println("Car model: ", v.model)
	case Bike:
		fmt.Println("Bike color: ", v.color)
	default:
		fmt.Println("Unexpected type")
	}
}
```

Using empty interfaces comes with the risk of losing compile-time type checking, making it harder to ensure type safety. Use type assertions or switches to safely retrieve the underlying type of a value stored in an empty interface.

### Best Practices: Keep Interfaces Small
Interfaces should be small to keep them simple (KISS), flexible to allow mixing and matching of features, and to be combined into larger interfaces for more complex interfaces.
```go
// Creating bigger interfaces with smaller interfaces
type Reader interface {
    Read()
}

type Writer interface {
    Write()
}

// A type that implements ReadWriter must satisfy both Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
```

Combining interfaces allows for a more modular approach where each type only needs to implement the methods that are relevant to its role. Also allows for keeping parts of a system independent, only interacting with the necessary interfaces required to do their jobs.

### Best Practices: Behavior Over Implementation
Interfaces are best used to describe what a program should do rather than *how* a program should do it.
```go
type Notifier interface {
	Notify(message string)
}
```
The `Notifier` interface in this case allows for a flexible interface for notification methods. It tells the program *what* to do, notify something, but now *how* to do it. This interface can be implemented by multiple notification methods, such as SMS, email, or push notifications.
```go
package main

import "fmt"

func main() {

	var notifier Notifier
	
	notifier = EmailNotifier{}
	notifier.Notify("Hello via Email!")
	
	notifier = SMSNotifier{}
	notifier.Notify("Hello via SMS!")
}

// Define behavior-focused and abstract interface
type Notifier interface {
	Notify(message string)
}

  

// Implement the interface in different ways
type EmailNotifier struct{}

// EmailNotifier type implementation of Notifier interface
func (e EmailNotifier) Notify(message string) {
	fmt.Println("Sending email with message:", message)
}

type SMSNotifier struct{}

// SMSNotifier type implementation of Notifier interface
func (s SMSNotifier) Notify(message string) {
	fmt.Println("Sending SMS with message:", message)
}
```