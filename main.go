package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/agungfir98/mini-redis/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	s, err := app.New(ctx)
	if err != nil {
		panic(err)
	}
	s.Run()
}
