package main

import "fmt"

func main() {

	// init array #1
	var arr_1 [10]int
	fmt.Println("arr_1:", arr_1)
	fmt.Printf("type of arr_1 %T\n", arr_1)
	for i := 0; i < len(arr_1); i++ {
		arr_1[i] = i
	}
	fmt.Println("arr_1 after loop:", arr_1)

	// init array #2
	var arr_2 = [3]string {"hello", "world", "yellow"}
	fmt.Println("arr_2:", arr_2)
	for i := 0; i < len(arr_2); i++ {
		arr_2[i] = "same"
	}
	fmt.Println("arr_2 after loop:", arr_2)

	// init array #3
	arr_3 := [4]float64 {1.5, 2.6, 3.14}
	fmt.Println(arr_3)

	// more dimensions
	var array2D = [2][3]int {{1, 2, 3}, {4, 5, 6}}
	fmt.Println(array2D)
	fmt.Printf("type of arr_1 %T\n", array2D)
	for i := 0; i < len(array2D); i++ {
		for j := 0; j < len(array2D[i]); j++ {
			array2D[i][j] = 7
		}
	}
	fmt.Println(array2D)
}
