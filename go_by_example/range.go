package main

import "fmt"

/*
	lesson: range
*/
func main() {

// range with slices
	s := []int{1, 2, 3}
	sum := 0
	for _, value := range s {
		sum += value
	}
	fmt.Println("Sum =", sum, "Slice:", s)
	for i, num := range s {
		if num == 2 {
			fmt.Println("num", num, "index", i)
		}
	}

	m := map[string]string{
		"one": "apple",
		"two": "bannana",
		"three": "orange",
	}

	for k, v := range m {
		fmt.Println(k, ":", v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}
	for i, c := range "hello" {
		fmt.Printf("%d %c\n", i, c)
	}
}
