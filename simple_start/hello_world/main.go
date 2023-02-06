// GO CLI
// go build -> compiles go source code files
// go run -> compiles and executes one or two files
// go fmt -> formats all the code in each file in the current directory
// go install -> compiles and "installs" a package
// go get -> downloads the raw source code of someone else's package
// go test -> runs any tests associated with the current project

// declare package name to which this file belongs
// packages types:
// Executable - generates a file that we run
// Reusable - code used as 'helpers'. Good place to put reusable logic
package main

// import fmt package (it is included in go stdlib)
import "fmt"

// declare 'main' function
func main() {
	fmt.Println("Hello world!")
}
