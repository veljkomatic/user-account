package app

import (
	"context"

	"github.com/pkg/errors"
)

// taskExecutionResult is a result of task execution.
type taskExecutionResult struct {
	err error
	// You can add more fields here to carry additional information about the task execution
}

// RunFunc is a function that will be executed by the task.
type RunFunc func(ctx context.Context) error

// Task is a single task that will be executed by the app.
type Task interface {
	// Name is the name of the task and must be unique for a single app.
	Name() string

	// Execute task. If ctx is canceled task should return ErrCanceled.
	// Execute function must not panic.
	Execute(ctx context.Context) <-chan taskExecutionResult
}

var _ Task = (*task)(nil)

type task struct {
	name    string
	runFunc RunFunc
}

func NewTask(name string, runFunc RunFunc) Task {
	return &task{
		name:    name,
		runFunc: runFunc,
	}
}

func (t *task) Name() string {
	return t.name
}

// Execute task. If ctx is canceled task should return ErrCanceled.
func (t *task) Execute(ctx context.Context) <-chan taskExecutionResult {
	resultChan := make(chan taskExecutionResult, 1)
	go func() {
		defer close(resultChan)
		defer func() {
			if r := recover(); r != nil {
				// Recover from the panic
				var err error
				switch x := r.(type) {
				case error:
					err = x
				case string:
					err = errors.New(x)
				default:
					err = errors.New("unknown panic")
				}
				resultChan <- taskExecutionResult{err: err}
			}
		}()
		err := t.runFunc(ctx)
		resultChan <- taskExecutionResult{err: err}
	}()
	return resultChan
}
