package tracing

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

// todo
func PropagateContext(ctx context.Context) context.Context {

	newCtx := context.Background()
	newCtx = opentracing.ContextWithSpan(newCtx, opentracing.SpanFromContext(ctx))
	return newCtx
}
