`go doc flag` for all available functions.

The `flag` package can be used to add command line arguments when the program is run, such as with `go run main.go`.

There are many `-Var` types that the flag package uses. Typically, you can just use `StringVar`, `IntVar`, `BoolVar`, `FloatVar`. These follow the format `func ...Var(p *Type, name string, value Type, usage string)` where `p *Type` is a pointer to the variable we want to change (so we can directly access it), the `name` of the flag that will be typed in the console, the default `value` of the flag, and the `usage` of the flag.  These functions *do not scan the command line*.

To scan the command line for these flags, we must follow our flag definitions with `flag.Parse()`.

For example:
```go
// main.go
package main

import (
	"flag"
	"fmt"
)

type language string

func main() {
	var lang string  // Declare variable that will hold the value of the flag
	flag.StringVar(&lang,  // Reference to lang variable
		"lang",  // Name of the flag when used in CLI
		"en",    // Default value of the flag if flag is not provided
		"The required language, e.g. en, fr, de...")  // Description/usage of the flag
	flag.Parse()  // Scan CLI for flags
	
	greeting := greet(language(lang)) // Parse "lang" from string to language type
	fmt.Println(greeting)
}

func greet(l language) string {
	...
}
```

In the above example, the `flag` package is imported. We are expecting a language input, so we assign a variable the closest type `string`. Then we create our flag with `flag.StringVar(...)`, making sure to pass a reference to our `string` variable `lang` and providing the name of the flag, default value, and usage of the flag. Finally, we use `flag.Parse()` to scan the command line for the flags. If the flags are not given when the program is run, they will use their default values assigned in `flag.StringVar(...)`.

There is also a shorthand for assigning variables to flags:
```go
exampleString := flag.String(name string, value string, usage string)
```
Note we are using a different function, `flag.String(...)` instead of `flag.StringVar(...)`, because it returns a pointer to the type so it can be assigned to the variable with the short declaration operator `:=`.