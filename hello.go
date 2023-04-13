package main

import "fmt"

func main() {
	fmt.Println("hello vscode!")
	a := foo()
	fmt.Println(a)

	b := bar(3)
	fmt.Println(b)
}

func foo() int {
	return 42
}

func bar(i int) int {
	return i + foo()
}
