package main

import (
	"context"
	"fmt"
	"time"
)

/*
	context.Background(): 根context
	context.WithDeadline: 派生context
*/

func main() {
	d := time.Now().Add(3 * time.Second)
	fmt.Println(time.Now())
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			fmt.Println(time.Now())
			return
		default:
			fmt.Println("1")
		}
	}
}
