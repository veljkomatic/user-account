package api

import "context"

// APIContext is the context for the api layer
// it can be expanded to include more information about the identity, project etc
type APIContext interface {
	Context() context.Context
}

type apiContext struct {
	ctx context.Context
}

func NewAPIContext(ctx context.Context) APIContext {
	return &apiContext{
		ctx: ctx,
	}
}

func (c *apiContext) Context() context.Context {
	return c.ctx
}
