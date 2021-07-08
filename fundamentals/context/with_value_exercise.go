package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "hello", "world")

	fmt.Println(ctx.Value("hello"))
}
