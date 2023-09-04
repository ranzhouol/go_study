package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	context.Background(): 根context
	context.WithCancel(): 派生context
*/

var wg sync.WaitGroup

func worker(ctx context.Context) {
	go worker2(ctx)
LOOP:
	for i := 0; ; i++ {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): //等待上级通知
			break LOOP
		default:
			fmt.Println("worker ", i)
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for i := 0; ; i++ {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): //等待上级通知
			break LOOP
		default:
			fmt.Println("worker2 ", i)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(3 * time.Second)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
