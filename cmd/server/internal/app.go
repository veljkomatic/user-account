package internal

import (
	"context"
	"net/http"

	"github.com/veljkomatic/user-account/cmd/server/internal/routers"
	"github.com/veljkomatic/user-account/common/app"
	"github.com/veljkomatic/user-account/common/cache/redis"
	"github.com/veljkomatic/user-account/common/log"
	"github.com/veljkomatic/user-account/common/server"
	"github.com/veljkomatic/user-account/common/storage/sql/postgres"
	"github.com/veljkomatic/user-account/pkg/user/api/handler"
	"github.com/veljkomatic/user-account/pkg/user/cache"
	"github.com/veljkomatic/user-account/pkg/user/repository"
	"github.com/veljkomatic/user-account/pkg/user/service"
)

// UserAccountServer is a struct that holds all the services that are used by the UserAccount API
type UserAccountServer struct {
	initialized       bool
	httpServerHandler http.Handler
	httpServer        *server.Server
	postgresClient    *postgres.Client
	redisClient       *redis.Client
}

// NewUserAccountServer creates a new UserAccountServer instance
func NewUserAccountServer() *UserAccountServer {
	return &UserAccountServer{
		initialized: false, // ensure that the server is not initialized when created
	}
}

// InitializeServer initializes UserAccountServer dependencies and configures the http server
func InitializeServer(srv *UserAccountServer, cfg *Config) app.CtxBuilderOption {
	return func(ctx context.Context, builder *app.AppBuilder) {
		defer func() { srv.initialized = true }()

		srv.redisClient = redis.MustConfigClient(ctx, cfg.Redis)
		srv.postgresClient = postgres.MustInitClient(ctx, cfg.Postgres)

		userHandler := initUserHandler(srv, cfg)
		srv.httpServerHandler = routers.InitRouter(userHandler)
		srv.httpServer = server.ConfigureServer(cfg.Server)

		builder.DeferCloser(srv.Shutdown)

		log.Info(
			ctx, "user account  server configured", log.Fields{
				"http": cfg.Server.Port,
			},
		)
	}
}

// initUserHandler initializes user handler and all the dependencies required for it
func initUserHandler(srv *UserAccountServer, cfg *Config) handler.UserHandler {
	var userHandler handler.UserHandler
	userRepository := repository.NewUserRepository(srv.postgresClient)
	userService := service.NewUserService(userRepository)

	if cfg.CacheEnabled {
		userCache := cache.NewUserCache(srv.redisClient)
		cachedUserService := service.NewCachedUserService(userService, userCache)
		userHandler = handler.NewUserHandler(cachedUserService)
		return userHandler
	}

	userHandler = handler.NewUserHandler(userService)
	return userHandler
}

// RegisterTasks registers tasks
// Currently only http server task is active
func (s *UserAccountServer) RegisterTasks(cfg *Config) app.CtxBuilderOption {
	return func(ctx context.Context, builder *app.AppBuilder) {
		if !s.initialized {
			panic("user account server not initialized")
		}

		startHttpServer := func(ctx context.Context) error {
			return s.httpServer.ListenAndServe(ctx, s.httpServerHandler)
		}
		builder.With(
			app.TaskDelegate("http-server", startHttpServer),
		)
	}
}

// Shutdown gracefully shuts down the server
func (s *UserAccountServer) Shutdown(ctx context.Context) {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Error(ctx, err, "http server shutdown failed")
	}
	if err := s.postgresClient.Close(); err != nil {
		log.Error(ctx, err, "postgres client shutdown failed")
	}
	if err := s.redisClient.Close(); err != nil {
		log.Error(ctx, err, "redis client shutdown failed")
	}
}
