package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
)

func initSignalHandlers(ctx context.Context, terminationCallback func()) <-chan struct{} {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	ch := make(chan struct{})

	go func() {
		<-ctx.Done()
		stop() // release resources, a second signal may interupt the callback execution
		defer close(ch)

		if terminationCallback != nil {
			terminationCallback()
		}
	}()

	return ch
}

func main() {
	done := initSignalHandlers(context.Background(), nil)
	<-done

	log.Println("done!")
}
