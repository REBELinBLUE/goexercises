package main

import "fmt"

func multiply(a int) int {
	return a * a
}

func multipleReturnTypes() (int, string) {
	return 5, "foo"
}

func nakedReturn() (isSomething bool) {
	isSomething = true
	return
}

func main() {
	// package name followed by function
	fmt.Println("Hello world!")
	fmt.Printf("Multiply - %v\n\n", multiply(5))

	a, b := multipleReturnTypes()
	fmt.Printf("Multiple Return Types a - %v\n", a)
	fmt.Printf("Multiple Return Types b - %v\n\n", b)

	fmt.Printf("Naked return - %v\n\n", nakedReturn())

}
