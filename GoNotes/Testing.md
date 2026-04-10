### Example Syntax
[Examples Docs](https://pkg.go.dev/testing#hdr-Examples)
 >The naming convention to declare examples for the package, a function F, a type T and method M on type T are:
>`func Example() { ... }`
>`func ExampleF() { ... }` 
>`func ExampleT() { ... }` 
>`func ExampleT_M() { ... }`
>
>Multiple example functions for a package/type/function/method may be provided by appending a distinct suffix to the name. The suffix must start with a lower-case letter.
>
>`func Example_suffix() { ... }`
>`func ExampleF_suffix() { ... }`
>`func ExampleT_suffix() { ... }`
>`func ExampleT_M_suffix() { ... }`

When testing a package, internal and external data can be tested regardless of how data is exposed. Private functions and public functions can both be tested in `*_test.go` packages. Test files are unique in that they implement a comment `// Output:` that the compiler uses to test output against. For example, if we had a package `main.go` that outputs `Hello, world!` in `main()`, we can use a `main_test.go` file to verify the output, even though `main()` is not a public function:

```go
// main.go
package main
import "fmt"
func main() {
	fmt.Println("Hello, world!")
}
```

```go
// main_test.go
package main

func Example_main() {
	main()
	// Output:
	// Hello, world!
}
```

> NOTE: This can be found on page 16 of **Learn Go with Pocket-Sized Projects**, however the Example function in the test file is wrong. See [this StackOverflow question](https://stackoverflow.com/q/79642686). Per the top solution, if we are testing main, we can simply add a suffix with a distinct name/identifier. *This identifier does not have to present in the main package.*

The test can be run with the command `go test` in the project directory.

```
learning-go/hello > go test
PASS
ok  	github.com/emersonds/learning-go/hello	0.001s
```