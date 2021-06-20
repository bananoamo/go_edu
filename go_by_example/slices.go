package main

import "fmt"

/*
	working with slices
*/
func main() {
	// slice declarations
	
	var firstSlice []int
	var	firstSliceDublicate []int

	fmt.Println("firstSlice :", firstSlice, "len is", len(firstSlice), "capacity is", cap(firstSlice))
	firstSlice = make([]int, 5, 10)
	fmt.Println("firstSlice :", firstSlice, "len is", len(firstSlice), "capacity is", cap(firstSlice))
	for i := 0; i < len(firstSlice); i++ {
		firstSlice[i] = i + 1
	}
	fmt.Println("firstSlice :", firstSlice, "len is", len(firstSlice), "capacity is", cap(firstSlice))

	/*
		slice capacity увеличивается в 2 раза если при добавлении в слайс привысить текущее значение
		slice len отображает текущее количество элементов в слайсе
	*/

	firstSlice = append(firstSlice, 6, 7, 8, 9, 10, 11)
	fmt.Println("firstSlice :", firstSlice, "len is", len(firstSlice), "capacity is", cap(firstSlice))

	firstSliceDublicate = make([]int, len(firstSlice))

	// copy makes copy of the slice by allocating new memory on heap i.e.
	// different pointers
	copy(firstSliceDublicate, firstSlice)
	fmt.Println("firstSliceDublicate :", firstSliceDublicate, "len is", len(firstSliceDublicate), "capacity is", cap(firstSliceDublicate))
	fmt.Printf("firstSlice adress: %p, firstSliceDublicate adress: %p\n", firstSlice, firstSliceDublicate)

	// just changing pointers, same pointers
	firstSliceDublicate = firstSlice
	fmt.Printf("firstSlice adress: %p, firstSliceDublicate adress: %p\n", firstSlice, firstSliceDublicate)
	fmt.Println(firstSlice, firstSliceDublicate)
	firstSlice[0] = 777
	fmt.Println(firstSlice, firstSliceDublicate)
	firstSliceDublicate = nil

	// working width slices of slices
	var stringSlice []string
	var	newStringSlice []string

	stringSlice = []string{"hello", "world", "yellow", "brother", "five"}
	fmt.Println(stringSlice)
	fmt.Printf("stringSlice type is : %T\n", stringSlice)

	newStringSlice = stringSlice[:]
	fmt.Println(newStringSlice)

	newStringSlice = newStringSlice[0:2]
	fmt.Println(newStringSlice)
	
	//dimention slices
	var array2D = make([][][]int, 5)
	fmt.Println(len(array2D))
	array2D = append(array2D, [][]int{[]int{1,2,3}})
	fmt.Println(len(array2D), cap(array2D), array2D)

	var zero []int
	fmt.Println(zero, len(zero), cap(zero), zero == nil)

	var arr = [5]int {1, 2, 3, 4, 5}
	partOfSlice := arr[1:4]
	fmt.Println(partOfSlice, len(partOfSlice), cap(partOfSlice))
	partOfSlice = partOfSlice[:cap(partOfSlice)]
	fmt.Println(partOfSlice, len(partOfSlice), cap(partOfSlice))

	// growing slices
	var zeroOne []int // nil

	for i := 1; i < 10; i++ {
		zeroOne = append(zeroOne, i)
	}
	var secondOne []int
	secondOne = append(secondOne, zeroOne...)
	fmt.Println(zeroOne)
	fmt.Println(secondOne)
	fmt.Printf("%p\n", zeroOne)
	fmt.Printf("%p\n", secondOne)




























}
