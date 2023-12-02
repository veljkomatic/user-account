package log

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/veljkomatic/user-account/common/environment"
	"github.com/veljkomatic/user-account/common/ptr"
)

const FieldLogScopeID string = "scopeId"

type scopeKey string

const ctxScopeKey scopeKey = "scopeKey"

type scope struct {
	id          string
	envAccessor environment.EnvAccessor
}

func newScope() *scope {
	return &scope{
		id:          uuid.New().String(),
		envAccessor: ptr.From(environment.RealEnvAccessor{}),
	}
}

func (s *scope) enrichLog(ctx context.Context, fields ...Fielder) []zap.Field {
	zapFields := s.parseFields(fields...)
	if !environment.Is(s.envAccessor, environment.EnvDevelopment) {
		zapFields = append(zapFields, zap.String(FieldLogScopeID, s.id))
	}
	return zapFields
}

func (s *scope) parseFields(fields ...Fielder) []zap.Field {
	if len(fields) == 0 {
		return []zap.Field{}
	}

	zapFields := make([]zap.Field, 0, len(fields[0].LogFields()))
	for _, f := range fields {
		for k, v := range f.LogFields() {
			if concrete, ok := v.(zap.Field); ok {
				zapFields = append(zapFields, concrete)
				continue
			}
			zapFields = append(zapFields, zap.Any(k, v))
		}
	}
	return zapFields
}

// ContextWithScope will return the context with a new scope added.
func ContextWithScope(ctx context.Context) context.Context {
	if val := ctx.Value(ctxScopeKey); val == nil {
		ctx = context.WithValue(ctx, ctxScopeKey, newScope())
	}
	return ctx
}
