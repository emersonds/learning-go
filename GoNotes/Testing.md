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

type language string

func main() {
	greeting := greet("en")
	fmt.Println("Hello, world!")
}

func greet(l language) string {
	switch l {
	case "en":
		fmt.Sprintf("Hello world")
	case "fr":
		fmt.Sprintf("Bonjour le monde")
	}
}
```

```go
// main_test.go
package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello world
}
```

> NOTE: This can be found on page 16 of **Learn Go with Pocket-Sized Projects**, however the Example function in the test file is wrong. See [this StackOverflow question](https://stackoverflow.com/q/79642686). Per the top solution, if we are testing main, we can simply add a suffix with a distinct name/identifier. *This identifier does not have to present in the main package.*

Breaking down this test, the Example function only checks the output of `main()`. As long as `main()` outputs "Hello world", it will pass. If `greet(l language)` is given a parameter `"fr"`, it will return "Bonjour le monde", which will fail this example because the example only accounts for the English input. Table-Driven tests can be used to check multiple different outputs.

These tests can be run with the command `go test` in the project directory.

```
learning-go/hello > go test
PASS
ok  	github.com/emersonds/learning-go/hello	0.001s
```

### Table-Driven Tests
Test functions can utilize tables to iterate over different test cases. This is table-driven tests. With this method, a table of test cases can be made by declaring a struct of different test cases with expected values. For example, if we want our "hello world" program to display different languages, we can check each output with a table using a struct with a `language` (custom string type) and a `want` (string):
```go
// main_test.go (continued)
...
func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	// Preparation phase, defines expected values
	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `unsupported language: "akk"`,
		},
		"Greek": {
			lang: "el",
			want: "Χαίρετε Κόσμε",
		},
		"Empty": {
			lang: "",
			want: `unsupported language: ""`,
		},
	}

	// Execution phase, Range over each language
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			// Decision phase, check returned value
			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
```

Let's break down this test function: 
- First, we are now calling it `TestGreet(t *testing.T)`. Rather than getting a single expected output **Example**, we are going to **Test** multiple test cases and their outputs.
- Next, we define a `testCase` struct that takes a `language` type (the input) and `want` string type, the expected output. Remember that the `language` type is a custom `string` type declared in `main.go`.
- After defining what a test case looks like, we create a test table `tests` that is a map of test cases and their expected outputs. This is called the ***preparation phase*** because we are declaring the test cases we will be testing our `greet()` function against. We use a map like a C# dictionary, where we are using a `string` name for the key, and a `testCase` struct for the value.
- Then, we `range` over the tests map, assigning the real output to a variable `got`. This is called the ***execution phase*** because we are executing the function and gathering the real output of each test case. Do note the function `t.Run(name, func(t *testing.T) {...}` in the range over `tests`. In this case, we call the name (`string` key of `tests`) of each `testCase` struct and get the real output before checking it against the expected output.
- Finally, we check the real output against the expected output to determine if our function is providing the correct output. This is called the ***decision phase*** because the test is deciding if the real output matches the expected output. This phase is simply checking if `got == tc.want`, and if it's not equal it returns an error message.