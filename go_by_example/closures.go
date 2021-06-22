package main

import "fmt"

//closures

func count() func() int {
	i := 0
	fmt.Println("outer var i", i)
	return func() int {
		fmt.Println("inner var i", i)
		i++
		return i
	}
}
func main() {
	nextInt := count()	
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := count()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
}
