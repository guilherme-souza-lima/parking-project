package main

import (
	"ProjetoEstacionamento/cmd"
	"ProjetoEstacionamento/infra"
	"context"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	config := infra.NewConfig()
	container := infra.NewContainer(config)
	cmd.StartHTTP(ctx, container)
}
