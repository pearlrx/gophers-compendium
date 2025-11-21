//# Hello, World in Go
//
//This is the traditional first program in Go.
//The goal of this topic is to:
//
//- Understand the minimal structure of a Go program
//- Learn how `package main` and `func main()` work
//- See how to print text to the console
//- Run a Go program using `go run`

//Minimal Go program

//Breakdown
//
//package main
//Entry point of an executable Go program. Any program that you want to run directly must have package main.
//
//import "fmt"
//Imports the standard library package fmt (formatted I/O). We use it for printing.
//
//func main()
//The main function. Execution starts here.
//
//fmt.Println("Hello, Go!")
//Calls the Println function from the fmt package to print a line to the console.

package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
