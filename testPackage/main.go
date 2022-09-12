package main

import (
	"errors"
	"fmt"
	"math"
	"storageModule/goroutine"
	"storageModule/shape"
	"storageModule/storage"
	"time"
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

func printSquare(shape shape.Shape) {
	square, _ := shape.GetSquare()
	fmt.Println(square)
}

func generateNumbers(n int, res chan int) {
	for i := 0; i < n; i++ {
		res <- i
	}
}

func main() {
	deadLock := make(chan int)

	go func() {
		deadLock <- 42
	}()

	time.Sleep(time.Microsecond * 1000)
	select {
	case n := <-deadLock:
		fmt.Println(n)
	default:
		fmt.Println("nothing")
	}

	t := time.Now()

	result1 := make(chan int)
	result2 := make(chan int)

	go func() {
		for {
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(time.Microsecond * 1000)
			}
		}
	}()
	go goroutine.CalcSomething(9000, result1)
	go goroutine.CalcSomething(4000, result2)
	fmt.Printf("result1: %d\n", <-result1)
	fmt.Printf("result2: %d\n", <-result2)
	time.Sleep(time.Second * 1)

	fmt.Printf("whole time: %s\n", time.Since(t))
	fmt.Println("--------")

	var square, triangle shape.Shape

	square = shape.NewSquare(10)
	triangle = shape.NewTriangle(10, 10, 30)
	printSquare(square)
	printSquare(triangle)
	square = triangle

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
