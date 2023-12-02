package app

import (
	"context"
	"github.com/veljkomatic/user-account/common/config"
	"github.com/veljkomatic/user-account/common/environment"
	"github.com/veljkomatic/user-account/common/log"
	"github.com/veljkomatic/user-account/common/ptr"
)

// AppBuilder is a builder for App.
type AppBuilder struct {
	appName string
	version string
	options []CtxBuilderOption

	environment environment.Environment
	envAccessor environment.EnvAccessor

	postStopHandlers []StopHandler
	closers          []CtxCloseFunc
	preStartHandlers []StartHandler
	tasks            map[string]Task
}

type AppBuilderOption func(builder *AppBuilder)

type CtxBuilderOption func(ctx context.Context, builder *AppBuilder)

func Builder(baseConfig config.BaseConfig) *AppBuilder {
	envAccessor := environment.RealEnvAccessor{}
	builder := &AppBuilder{
		appName:     environment.ServiceName(ptr.From(envAccessor)),
		version:     environment.ServiceVersion(ptr.From(envAccessor)),
		environment: environment.Get(ptr.From(envAccessor)),
		tasks:       make(map[string]Task),
		envAccessor: ptr.From(envAccessor),
	}

	builder = builder.With(VersionFromEnv(ptr.From(envAccessor)))

	return builder
}

func (ab *AppBuilder) Name() string {
	return ab.appName
}

func (ab *AppBuilder) With(options ...AppBuilderOption) *AppBuilder {
	ab.options = append(ab.options, ab.wrapOptions(options...)...)
	return ab
}

func (ab *AppBuilder) addPostStopHandler(handler StopHandler) *AppBuilder {
	ab.postStopHandlers = append(ab.postStopHandlers, handler)
	return ab
}

func (ab *AppBuilder) addPreStartHandler(handler StartHandler) *AppBuilder {
	ab.preStartHandlers = append(ab.preStartHandlers, handler)
	return ab
}

func (ab *AppBuilder) DeferCloser(closer CtxCloseFunc) *AppBuilder {
	ab.closers = append(ab.closers, closer)
	return ab
}

func (ab *AppBuilder) WithCtx(option ...CtxBuilderOption) *AppBuilder {
	ab.options = append(ab.options, option...)
	return ab
}

func (ab *AppBuilder) getAppName(ctx context.Context) string {
	if ab.appName != "" {
		return ab.appName
	}

	if environment.Is(ab.envAccessor, environment.EnvDevelopment) {
		log.Warn(
			ctx,
			"App name not provided, using default app name. Provide app name or set fallback app name to avoid this warning.",
		)
		return "default-app-name"
	}

	return ""
}

func (ab *AppBuilder) Build(ctx context.Context) *App {
	ctx, cancel := context.WithCancelCause(ctx)
	for i := 0; i < len(ab.options); i++ {
		ab.options[i](ctx, ab)
	}

	appName := ab.getAppName(ctx)
	if appName == "" {
		panic("app name must be non empty")
	}

	var tasks []Task
	for _, task := range ab.tasks {
		tasks = append(tasks, task)
	}

	return newApp(
		ctx,
		cancel,
		appName,
		ab.version,
		ab.environment,
		tasks,
		ab.postStopHandlers,
		ab.closers,
		ab.preStartHandlers,
	)
}

func (ab *AppBuilder) SetVersion(version string) *AppBuilder {
	if version != "" {
		ab.version = version
	}

	return ab
}

func (ab *AppBuilder) SetEnv(env environment.Environment) *AppBuilder {
	ab.environment = env
	return ab
}

func (ab *AppBuilder) setName(name string) *AppBuilder {
	ab.appName = name
	return ab
}

func (ab *AppBuilder) addTask(taskName string, runFunc RunFunc) *AppBuilder {
	newTask := NewTask(taskName, runFunc)
	_, duplicate := ab.tasks[newTask.Name()]
	if duplicate {
		log.Debug(
			context.TODO(), "task with same name already added", log.Fields{
				"task": newTask.Name(),
			},
		)
		return ab
	}
	ab.tasks[newTask.Name()] = newTask
	return ab
}

func (ab *AppBuilder) wrapOption(option AppBuilderOption) CtxBuilderOption {
	return func(ctx context.Context, builder *AppBuilder) {
		option(builder)
	}
}

func (ab *AppBuilder) wrapOptions(options ...AppBuilderOption) []CtxBuilderOption {
	var ctxOptions []CtxBuilderOption
	for _, option := range options {
		ctxOptions = append(ctxOptions, ab.wrapOption(option))
	}

	return ctxOptions
}
