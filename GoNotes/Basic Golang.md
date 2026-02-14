`.go` files can be organized into packages and imported to another file. For example, writing a calculator program, files for calculation could be stored in `package calc` and files for input/output stored in `package io`.

A **package declaration** tells the Go compiler which package this file belongs to, where `package main` will be compiled into an executable on `go build main`.

An **import statement** allows the use of code from other packages: `import "fmt"`

`func` declares a function:
```go
func main () {
	fmt.Println("Hello World")
}
```

`.go` files can be built and run using the commands `go build main.go` and `./main`, where `main` is the resulting executable file created by `go build`. A `.go` file can also be run using `go run main.go` but the `go run` command **does not build an executable file**.
### Go Doc
`go doc` is used to extract and view documentation from `.go` files. View package docs by running `go doc <package name>` or info about a package function by running `go doc <package name>.<function>`

`go doc fmt` or `go doc fmt.Println`

[Go's official documentation](https://go.dev/doc/)

### Example function to print an ASCII dog
```go
package main

import f "fmt"

func main () {
	f.Printf(" __ _\n" +

		"o'')}____//\n" +

		" `_/ )\n" +

		" (_(_/-(_/\n")

}
```
