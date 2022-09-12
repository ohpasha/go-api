package goroutine

import (
	"fmt"
	"time"
)

func CalcSomething(n int, res chan int) {
	t := time.Now()
	result := 0

	for i := 0; i <= n; i++ {
		result = i * 3
		time.Sleep(time.Microsecond * 3)
	}

	fmt.Printf("difficult result: %d, time: %s\n", result, time.Since(t))
	res <- result
}
