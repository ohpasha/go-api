package main

import (
	"errors"
	"fmt"
	"math"
	"storageModule/storage"
)

func calcCircleArea(r int) (float64, error) {
	if r < 0 {
		return float64(0), errors.New("r не можут быть меньше нуля")
	}
	return math.Pi * math.Pow(float64(r), 2), nil

}

func printCircleArea(r int) {
	circleArea, error := calcCircleArea(r)

	if error != nil {
		fmt.Println(error.Error())
		return
	}

	fmt.Printf("circle Area is %f64 with radius %d", circleArea, r)
}

func reqReverse(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	return append(reqReverse(arr[1:]), arr[0])
}

func newEmployee(name string, age int) storage.Employee {
	return storage.Employee{
		Name: name,
		Age:  age,
	}
}

func checkType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	}
}

func main() {
	checkType("str")
	var s storage.Storage
	fmt.Printf("type s %T\n", s)
	s = storage.NewMemoryStorage()
	fmt.Printf("type s %T\n", s)

	var array = []int{1, 2, 3, 4}
	var b = []int{5, 5, 5}
	var currentMap = make(map[string]string)

	currentMap["asdasd"] = "dd"

	g := newEmployee("name", 22)
	fmt.Printf("%+v\n", g)

	a := append(array, b...)
	fmt.Println(a)

	fmt.Println(reqReverse(array))
	g.SetName("g")
	fmt.Println(g.GetInfo())

	printCircleArea(2)
}
