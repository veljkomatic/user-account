package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"

	"github.com/veljkomatic/user-account/cmd/command/internal"
	"github.com/veljkomatic/user-account/cmd/command/internal/database"
	"github.com/veljkomatic/user-account/common/config"
	"github.com/veljkomatic/user-account/common/environment"
	"github.com/veljkomatic/user-account/common/log"
	"github.com/veljkomatic/user-account/common/ptr"
)

func main() {
	ctx := context.Background()
	cfg := config.Builder[*internal.CommandConfig]().Build(ctx)

	log.Info(ctx, "config loaded")

	ctx, cancel := context.WithCancel(context.Background())
	go overloadInterrupt(cancel)

	envAccessor := environment.RealEnvAccessor{}
	cmd := cli.NewApp()
	cmd.EnableBashCompletion = true
	cmd.Name = "command"
	cmd.Usage = "Internal tooling."
	cmd.Description = description
	cmd.Commands = []*cli.Command{
		database.NewCmd(cfg, ptr.From(envAccessor)),
	}

	if err := cmd.RunContext(ctx, os.Args); err != nil {
		log.Error(ctx, err, "failed running command")
		os.Exit(1)
	}
}

func overloadInterrupt(fn func()) {
	sigint := make(chan os.Signal, 1)
	// interrupt signal sent from terminal
	signal.Notify(
		sigint,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	<-sigint

	fn()
}

const description = `command is an internal tool for the platform.
Used as an umbrella for all scripts, commands and jobs needed inside the engineering team.
Everything that is possible should be always done with command as we have bigger control and more flexibility.`
