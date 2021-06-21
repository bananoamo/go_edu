package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// if args is the same type
func multiple(a, b int) int {
	return a * b
}

// for multiple return value
func vals(a, b int) (int, int) {
	return a, b
}

// variadic function
func sum(nums ...int){
	fmt.Print(nums, " ")
	sum := 0
	for _, vals := range nums {
		sum += vals
	}
	fmt.Println(sum)
}


func main() {
	fmt.Println(plus(1, 2))
	fmt.Println(multiple(2, 2))

	// multiple return values
	_, z := vals(99, 66)
	fmt.Println(z)

	// variable number of arguments
	sum(1, 2, 3)
	sum([]int{5,6,7}...)
	var slice = []int{10, 11, 12}
	sum(slice...)

}

