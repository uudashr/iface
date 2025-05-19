package tracing

import "context"

type Tracer interface{}

var GlobalTracer Tracer

type Span interface {
	Finish()
}

type span struct{}

func (s *span) Finish() {
}

func StartSpanFromContext(ctx context.Context) (Span, error) {
	return StartSpanFromContextWithTracer(ctx, GlobalTracer)
}

func StartSpanFromContextWithTracer(ctx context.Context, tracer Tracer) (Span, error) {
	return nil, nil
}

func StartSpanFromContext2(ctx context.Context) (Span, error) { // want "'StartSpanFromContext2' function return 'Span' interface at the 1st result, abstract a single concrete implementation of '\\*span'"
	return startSpanFromContextWithTracer(ctx, GlobalTracer)
}

func startSpanFromContextWithTracer(ctx context.Context, tracer Tracer) (*span, error) {
	return nil, nil
}
