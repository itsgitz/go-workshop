package main

import (
	"context"
	"fmt"
	"time"
)

// this is one millisecond (1ms)
const shortDuration = 1 * time.Millisecond

func main() {
	d := time.Now().Add(shortDuration)

	fmt.Println("shortDuration:", shortDuration)
	fmt.Println("d:", d)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
