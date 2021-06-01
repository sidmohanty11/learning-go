package main //package -> project/workspace
//two types of packages ->
// 1. <!executable type!> -> generates the file that we run
// 2. <!reusable type!> -> code used as helpers (reusable logic) deps/libs/helperfuncs

import "fmt" //format lib(standard), get access(link) to the lib!

func main() { //just a function
	fmt.Println("Hello World")
}

//go run ... (compiles and executes)
//go build ... (only compiles -> hello.exe)
//go fmt ... (formats all code)
//go install ... (handle dependencies)
//go get ... (handle dependencies)
//go test ... (runs any test)
