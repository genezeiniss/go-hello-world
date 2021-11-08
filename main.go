/*
go have two types of packages:
	1. executable -> generates executable file (the file that we can run)
	2. reusable -> code used as "utils" or "helpers"

package main -> is a secret that means that class has runnable code
main package must have a function "man" as well
*/
package main

/*
fmt -> format -> standard library package to print different information
for other standard libraries: https://pkg.go.dev/std
*/
import "fmt"

/*
to run the project:
	open terminal -> navigate to hello-word project
	$ go run [file you want to execute]
	example: go run main.go
*/
func main() {
	fmt.Println("Hi, there!")
}

/*
go commands:
	$ go build main.go -> just compiles the file without to execute it.
		it creates the executable file (main) without to run it.
		to execute this file, run $ ./main
	$ go run main.go -> runs one or two files
	$ go ftm -> formats all the code in each file in the current directory
*/
