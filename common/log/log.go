/*
Package log contains the logger interface and it's nil implementation
Log levels / verbosity go from lowest to highest: Debug, Info, Warn, Error.
Setting the logger verbosity to a value will print that or higher level logs.
*/
package log

import (
	"context"
)

type Level = string

const (
	// LevelDebug is used for printing debug logs and higher.
	LevelDebug Level = "debug"
	// LevelInfo is used for printing info logs and higher.
	LevelInfo Level = "info"
	// LevelWarn is used for printing warning logs and higher.
	LevelWarn Level = "warning"
	// LevelError is used for printing error logs only.
	LevelError Level = "error"
)

var global Logger = NewZapLogger(&Config{})

// Logger is an interface which declares logging with context
// Logging will  try to extract the current logging scope and group the logs.
type Logger interface {
	Debug(ctx context.Context, msg string, fields ...Fielder)
	Info(ctx context.Context, msg string, fields ...Fielder)
	Warn(ctx context.Context, msg string, fields ...Fielder)
	Error(ctx context.Context, err error, msg string, fields ...Fielder)
}

// Debug is a wrapper around global's log Debug
func Debug(ctx context.Context, msg string, fields ...Fielder) {
	global.Debug(ctx, msg, fields...)
}

// Info is a wrapper around global's log Info
func Info(ctx context.Context, msg string, fields ...Fielder) {
	global.Info(ctx, msg, fields...)
}

// Warn is a wrapper around global's log Warn
func Warn(ctx context.Context, msg string, fields ...Fielder) {
	global.Warn(ctx, msg, fields...)
}

// Error is a wrapper around global's log Error
func Error(ctx context.Context, err error, msg string, fields ...Fielder) {
	global.Error(ctx, err, msg, fields...)
}
