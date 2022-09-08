package module

import (
	"context"
	"fmt"
	"github.com/zhashkevych/scheduler"
	"time"
)

func main() {
	s := scheduler.NewScheduler()

	s.Add(context.Background(), func(ctx context.Context) {
		fmt.Printf("time %s\n", time.Now())
	}, time.Second)

	time.Sleep(time.Minute)
}
