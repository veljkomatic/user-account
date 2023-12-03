package app

import (
	"context"
	"github.com/veljkomatic/user-account/common/config"
	"github.com/veljkomatic/user-account/common/environment"
	"github.com/veljkomatic/user-account/common/log"
	"github.com/veljkomatic/user-account/common/ptr"
)

// Builder is a builder for App.
type Builder struct {
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

type BuilderOption func(builder *Builder)

type CtxBuilderOption func(ctx context.Context, builder *Builder)

// NewBuilder creates new app builder.
// currently we do not use config.BaseConfig, but we will use it in the future.
func NewBuilder(baseConfig config.BaseConfig) *Builder {
	envAccessor := environment.RealEnvAccessor{}
	builder := &Builder{
		appName:     environment.ServiceName(ptr.From(envAccessor)),
		version:     environment.ServiceVersion(ptr.From(envAccessor)),
		environment: environment.Get(ptr.From(envAccessor)),
		tasks:       make(map[string]Task),
		envAccessor: ptr.From(envAccessor),
	}

	builder = builder.With(VersionFromEnv(ptr.From(envAccessor)))

	return builder
}

// Name gets app name.
func (b *Builder) Name() string {
	return b.appName
}

// DeferCloser adds CtxCloseFunc to the app builder.
func (b *Builder) DeferCloser(closer CtxCloseFunc) *Builder {
	b.closers = append(b.closers, closer)
	return b
}

// With adds BuilderOption to the app builder.
func (b *Builder) With(options ...BuilderOption) *Builder {
	b.options = append(b.options, b.wrapOptions(options...)...)
	return b
}

// WithCtx adds CtxBuilderOption to the app builder.
func (b *Builder) WithCtx(option ...CtxBuilderOption) *Builder {
	b.options = append(b.options, option...)
	return b
}

// addPostStopHandler adds StopHandler to the app builder.
func (b *Builder) addPostStopHandler(handler StopHandler) *Builder {
	b.postStopHandlers = append(b.postStopHandlers, handler)
	return b
}

// addPreStartHandler adds StartHandler to the app builder.
func (b *Builder) addPreStartHandler(handler StartHandler) *Builder {
	b.preStartHandlers = append(b.preStartHandlers, handler)
	return b
}

// getAppName returns app name.
func (b *Builder) getAppName(ctx context.Context) string {
	if b.appName != "" {
		return b.appName
	}

	if environment.Is(b.envAccessor, environment.EnvDevelopment) {
		log.Warn(
			ctx,
			"App name not provided, using default app name. Provide app name or set fallback app name to avoid this warning.",
		)
		return "default-app-name"
	}

	return ""
}

// Build builds App.
func (b *Builder) Build(ctx context.Context) *App {
	ctx, cancel := context.WithCancelCause(ctx)
	for i := 0; i < len(b.options); i++ {
		option := b.options[i]
		option(ctx, b)
	}

	appName := b.getAppName(ctx)
	if appName == "" {
		panic("app name must be non empty")
	}

	var tasks []Task
	for _, task := range b.tasks {
		tasks = append(tasks, task)
	}

	return newApp(
		ctx,
		cancel,
		appName,
		b.version,
		b.environment,
		tasks,
		b.postStopHandlers,
		b.closers,
		b.preStartHandlers,
	)
}

// SetVersion sets app version.
func (b *Builder) SetVersion(version string) *Builder {
	if version != "" {
		b.version = version
	}

	return b
}

// SetEnv sets app environment.
func (b *Builder) SetEnv(env environment.Environment) *Builder {
	b.environment = env
	return b
}

// setName sets app name.
func (b *Builder) setName(name string) *Builder {
	b.appName = name
	return b
}

// addTask adds task to the app.
func (b *Builder) addTask(taskName string, runFunc RunFunc) *Builder {
	newTask := NewTask(taskName, runFunc)
	_, duplicate := b.tasks[newTask.Name()]
	if duplicate {
		log.Debug(
			context.TODO(), "task with same name already added", log.Fields{
				"task": newTask.Name(),
			},
		)
		return b
	}
	b.tasks[newTask.Name()] = newTask
	return b
}

// wrapOption wraps BuilderOption into CtxBuilderOption
func (b *Builder) wrapOption(option BuilderOption) CtxBuilderOption {
	return func(ctx context.Context, builder *Builder) {
		option(builder)
	}
}

// wrapOptions wraps BuilderOption into CtxBuilderOption
func (b *Builder) wrapOptions(options ...BuilderOption) []CtxBuilderOption {
	var ctxOptions []CtxBuilderOption
	for _, option := range options {
		ctxOptions = append(ctxOptions, b.wrapOption(option))
	}

	return ctxOptions
}
