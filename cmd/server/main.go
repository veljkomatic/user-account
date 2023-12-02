package main

import (
	"context"

	"github.com/veljkomatic/user-account/cmd/server/internal"
	"github.com/veljkomatic/user-account/common/app"
	"github.com/veljkomatic/user-account/common/config"
	"github.com/veljkomatic/user-account/common/environment"
	"github.com/veljkomatic/user-account/common/log"
	"github.com/veljkomatic/user-account/common/ptr"
)

func main() {
	ctx := context.Background()
	cfg := config.Builder[*internal.Config]().Build(ctx)

	log.Info(ctx, "config loaded")

	server := internal.NewUserAccountServer()
	envAccessor := environment.RealEnvAccessor{}

	fullServiceName := environment.FullServiceName(ptr.From(envAccessor))
	app.Builder(cfg).
		With(app.Name(fullServiceName, "user-account-server")).
		WithCtx(internal.InitializeServer(server, cfg)).
		WithCtx(server.RegisterTasks(cfg)).
		Build(ctx).
		Start(ctx)
}
