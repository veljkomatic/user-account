package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pkg/errors"

	"github.com/veljkomatic/user-account/common/environment"
	"github.com/veljkomatic/user-account/common/log"
)

// StopHandler represents operations that will be executed after the app is stopped.
// StopHandler must be non-blocking and short living func.
type StopHandler func(app *App)

// StartHandler represents operations that will be executed before the app start.
// StartHandler must be non-blocking and short living func.
type StartHandler func(app *App)

// CtxCloseFunc represents operations that will be executed after the app is stopped.
type CtxCloseFunc func(ctx context.Context)

// App is container for tasks execution.
// App is responsible for starting and stopping tasks.
type App struct {
	ctx    context.Context
	cancel context.CancelCauseFunc
	wg     sync.WaitGroup

	name        string
	version     string
	environment environment.Environment

	tasks           []Task
	postStopHandler []StopHandler
	closers         []CtxCloseFunc
	preStartHandler []StartHandler
}

func newApp(
	ctx context.Context,
	cancelFunc context.CancelCauseFunc,
	name string,
	version string,
	environment environment.Environment,
	tasks []Task,
	postStopHandler []StopHandler,
	closers []CtxCloseFunc,
	preStartHandler []StartHandler,
) *App {
	return &App{
		ctx:             ctx,
		cancel:          cancelFunc,
		name:            name,
		version:         version,
		environment:     environment,
		tasks:           tasks,
		closers:         closers,
		preStartHandler: preStartHandler,
		postStopHandler: postStopHandler,
	}
}

func (app *App) Name() string {
	return app.name
}

func (app *App) Version() string {
	return app.version
}

func (app *App) Environment() environment.Environment {
	return app.environment
}

// Start starts the app and blocks until the app is stopped.
func (app *App) Start(ctx context.Context) {
	log.Info(ctx, "starting app", log.Fields{"app_name": app.Name(), "app_version": app.Version()})
	app.registerCancellationOnInterrupt(app.cancel)

	app.preStart()
	app.startTasksExecution(app.ctx)
	app.wg.Wait() // Wait for all tasks to complete
	app.postStop()

	defer app.cancel(nil)
	log.Info(app.ctx, "app stopped")
}

// startTasksExecution starts all tasks in separate goroutines
func (app *App) startTasksExecution(
	appCtx context.Context,
) {
	for _, t := range app.tasks {
		app.wg.Add(1)
		app.startTaskExecutionAsync(appCtx, t)
	}
}

// startTaskExecutionAsync starts task execution in a separate goroutine
func (app *App) startTaskExecutionAsync(
	ctx context.Context,
	task Task,
) {
	go func() {
		defer app.wg.Done()
		log.Info(ctx, "starting task execution")

		resultChan := task.Execute(ctx)

		select {
		case <-ctx.Done():
			// Context was cancelled, handle graceful shutdown
			log.Info(ctx, "Context cancelled, shutting down task", log.Fields{"name": task.Name()})
			// Additional graceful shutdown logic can be placed here if needed
		case result := <-resultChan:
			// Task completed, handle the result
			if result.err != nil {
				log.Error(ctx, result.err, "Task execution error", log.Fields{"name": task.Name(), "error": result.err})
			}
			// Handle other aspects of the result if necessary
		}
	}()
}

// registerCancellationOnInterrupt registers cancellation on interrupt signal
func (app *App) registerCancellationOnInterrupt(cancelFunc context.CancelCauseFunc) {
	go func() {
		sigint := make(chan os.Signal, 1)
		// interrupt signal sent from terminal
		signal.Notify(
			sigint,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGHUP,
		)

		signal := <-sigint
		log.Warn(app.ctx, "sigint detected, cancelling app", log.Fields{"signal": signal})
		cancelFunc(errors.New("OSInterrupt"))
	}()
}

// preStart executes all pre-start handlers
func (app *App) preStart() {
	log.Info(app.ctx, "running app on-start hook")
	for _, startHandler := range app.preStartHandler {
		startHandler(app)
	}
}

// postStop executes all post-stop handlers and closers in reverse order
func (app *App) postStop() {
	log.Info(app.ctx, "running app on-stop hook")
	for _, stopHandler := range app.postStopHandler {
		stopHandler(app)
	}
	log.Info(app.ctx, "running closers in reverse order to mimic defer behaviour")
	for i := len(app.closers) - 1; i >= 0; i-- {
		closer := app.closers[i]
		closer(app.ctx)
	}
}
