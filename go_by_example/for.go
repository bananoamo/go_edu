package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Printf("\n")

	fmt.Println("j loop is here")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	fmt.Println("forever loop")
	for {
		fmt.Println("unlimited loop")
		break;
	}
}
