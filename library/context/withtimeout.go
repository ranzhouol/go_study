package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func worker3(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg2.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg2.Add(1)
	go worker3(ctx)
	time.Sleep(time.Second * 5) //如果修改为20毫秒，就会提前关闭goroutine
	cancel()                    // 通知子goroutine结束
	wg2.Wait()
	fmt.Println("over")
}
