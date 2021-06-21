package main

import "fmt"

func main() {
	
	var array = make(map[string]string)
	fmt.Printf("type: %T\n", array)

	array["k1"] = "one" 
	array["k2"] = "two"
	array["k4"] = ""
	fmt.Println(array)
	fmt.Println("Get k1:", array["k1"])
	fmt.Println("Get k2:", array["k2"])
	fmt.Println("map len:", len(array))
	val, err := array["k4"]
	fmt.Printf("%T, %T\n", val, err)
	fmt.Println(val == "", err)
	delete(array, "k1")
	fmt.Println(array)

	var m map[string]int

	m = make(map[string]int)
	m["route"] = 66
	i := m["route"]
	fmt.Println(m, i)
	fmt.Println("m[\"root\"]", m["root"])
	fmt.Println("len is:", len(m))
	delete(m, "root")

	_, ok := m["route"]
	fmt.Println(ok)
	for key, value := range array {
		fmt.Println("Key:", key, "value:", value)
	}

	commits := map[string]int {
		"age": 31,
		"length": 175,
		"class": 11,
	}
	fmt.Println(commits)

	newM := map[string]int{} // the same as newM := make(map[string]int)
	fmt.Println("newM :", newM)

	second := map[string]map[string]int{
		"hello":{
			"world":100
		}
	}
	fmt.Println(second)
	fmt.Println(second["hello"]["world"])
}
