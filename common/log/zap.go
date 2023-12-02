package log

import (
	"context"
	"os"

	"go.ajitem.com/zapdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/veljkomatic/user-account/common/environment"
	"github.com/veljkomatic/user-account/common/ptr"
)

var _ Logger = (*zapLogger)(nil)

// zapLogger is the log.Logger implementation.
// It contains a zap sugared logger.
type zapLogger struct {
	logger      *zap.Logger
	envAccessor environment.EnvAccessor
}

// NewZapLogger parses the config and returns a new logger.
func NewZapLogger(cfg *Config) *zapLogger {
	logger := &zapLogger{
		logger:      newLocalDevelopmentLogger(cfg),
		envAccessor: ptr.From(environment.RealEnvAccessor{}),
	}

	if environment.Is(logger.envAccessor, environment.EnvDevelopment) {
		return logger
	}

	productionLogger, err := newProductionDevelopmentLogger(cfg)
	if err != nil {
		return logger
	}
	logger.logger = productionLogger
	return logger
}

func newLocalDevelopmentLogger(cfg *Config) *zap.Logger {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.ConsoleSeparator = "  "
	jsonEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	atom := zap.NewAtomicLevel()
	atom.SetLevel(mapVerbosityLevel(cfg.Level))
	core := zapcore.NewCore(jsonEncoder, zapcore.Lock(os.Stdout), atom)

	logger := zap.New(core).WithOptions(
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.AddCaller(),
	)

	defer logger.Sync()
	return logger
}

func newProductionDevelopmentLogger(cfg *Config) (*zap.Logger, error) {
	encoderConfig := zapdriver.NewProductionEncoderConfig()
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	atom := zap.NewAtomicLevel()
	atom.SetLevel(mapVerbosityLevel(cfg.Level))
	core := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), atom)

	return zap.New(core).WithOptions(
		zapdriver.WrapCore(
			zapdriver.ReportAllErrors(true),
			zapdriver.ServiceName(cfg.ServiceName),
			zapdriver.ServiceVersion(cfg.Version),
		),
		zap.AddStacktrace(zap.ErrorLevel),
	), nil
}

// Debug prints debug logs. Only seen if verbosity level log.Debug.
// It will try to extract the scope from context before logging.
func (log *zapLogger) Debug(ctx context.Context, msg string, fields ...Fielder) {
	log.logger.Debug(msg, log.withContextualLogging(ctx, fields...)...)
}

// Info prints info logs. Only seen if verbosity level log.Info or lower.
// It will try to extract the scope from context before logging.
func (log *zapLogger) Info(ctx context.Context, msg string, fields ...Fielder) {
	log.logger.Info(msg, log.withContextualLogging(ctx, fields...)...)
}

// Warn prints warning logs. Only seen if verbosity level log.Warn or lower.
// It will try to extract the scope from context before logging.
func (log *zapLogger) Warn(ctx context.Context, msg string, fields ...Fielder) {
	log.logger.Warn(msg, log.withContextualLogging(ctx, fields...)...)
}

// Error prints error logs. Only seen if verbosity level log.Error or lower.
// It will try to extract the scope from context before logging.
func (log *zapLogger) Error(ctx context.Context, err error, msg string, fields ...Fielder) {
	fields = withError(err, fields...)
	log.logger.Error(msg, log.withContextualLogging(ctx, fields...)...)
}

func (log *zapLogger) withContextualLogging(ctx context.Context, fields ...Fielder) []zap.Field {
	ctx = ContextWithScope(ctx)
	scope, _ := ctx.Value(ctxScopeKey).(*scope)
	return scope.enrichLog(ctx, fields...)
}

func withError(err error, fields ...Fielder) []Fielder {
	fields = append(fields, Fields{"error": err})
	return fields
}

func mapVerbosityLevel(verbosity string) zapcore.Level {
	switch verbosity {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelError:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
