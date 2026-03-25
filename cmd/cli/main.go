package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/atlazar/visual-concurrency/internal/cli"
)

const (
	hangSec = 5
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	r := cli.NewApp(ctx)
	wg := r.Run()

	go func() {
		//If all goroutine complete - cancel context to gracefully exit process
		wg.Wait()
		cancel()
	}()

	<-ctx.Done()
	// Restore system signal handling in case of goroutine hang
	// Next SIGINT will result force exit
	cancel()
	wg.Wait()
	fmt.Println("Exit")
}
