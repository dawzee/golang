package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	if err := exec.CommandContext(ctx, "delay", "1").Run(); err != nil {
		fmt.Println("Run Error:", err)
	}

	fmt.Println("Run Complete:", ctx.Err())
	// cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("still waiting after 3 seconds!")
	case <-ctx.Done():
		fmt.Println("Done:", ctx.Err())
	}
}
