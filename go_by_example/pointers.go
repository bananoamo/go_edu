package main

import "fmt"

func simpleFunction(n int) {
	n += 1024
	fmt.Println("n inside simpleFunction", n)
	fmt.Println("n ptr in simpleFunction", &n)
}

func ptrFunction(n *int) {
	*n += 1024
	fmt.Println("n inside ptrFunction", *n)
	fmt.Println("n ptr in ptrFunction", n)
}

func main() {
	i := 1
	simpleFunction(i)
	fmt.Println("i outside simpleFunction", i)

	ptrFunction(&i)
	fmt.Println("i outside ptrFunction", i)

	fmt.Println(&i)
}
