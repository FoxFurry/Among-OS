package main

import (
	"context"
	"github.com/foxfurry/university/sommip/internal/boot"
	"github.com/foxfurry/university/sommip/internal/console/terminal"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	term := terminal.NewTerminal()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	boot.StartScreen()

	go term.Start(ctx)

	<-sigs

	cancel()
}

