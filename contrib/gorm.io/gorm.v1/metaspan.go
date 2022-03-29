package gorm

import (
	"math/rand"
	"time"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

var _ ddtrace.Tracer = (*NoopTracer)(nil)

// NoopTracer is an implementation of ddtrace.Tracer that is a no-op.
type NoopTracer struct{}

// StartSpan implements ddtrace.Tracer.
func (NoopTracer) StartSpan(operationName string, opts ...ddtrace.StartSpanOption) ddtrace.Span {
	return metaSpan{}
}

// SetServiceInfo implements ddtrace.Tracer.
func (NoopTracer) SetServiceInfo(name, app, appType string) {}

// Extract implements ddtrace.Tracer.
func (NoopTracer) Extract(carrier interface{}) (ddtrace.SpanContext, error) { // TODO implement ?
	return metaSpanContext{}, nil
}

// Inject implements ddtrace.Tracer.
func (NoopTracer) Inject(context ddtrace.SpanContext, carrier interface{}) error { return nil } // TODO implement ?

// Stop implements ddtrace.Tracer.
func (NoopTracer) Stop() {}

var _ ddtrace.Span = (*metaSpan)(nil)
var random *rand.Rand

// metaSpan is an implementation of ddtrace.Span that is a no-op.
type metaSpan struct {
	spanContext metaSpanContext

	startTime time.Time
	parentCtx ddtrace.SpanContext
}

func newMetaSpan(parentCtx ddtrace.SpanContext) metaSpan {
	return metaSpan{
		spanContext: metaSpanContext{
			traceID: parentCtx.TraceID(),
			spanID:  random.Uint64(),
		},
		startTime: time.Now(),
		parentCtx: parentCtx,
	}
}

// GetSpanID get the span ID
func (s metaSpan) GetSpanID() uint64 {
	return s.spanContext.SpanID()
}

// GetParentCtx get parents Span Context
func (s metaSpan) GetParentCtx() ddtrace.SpanContext {
	return s.parentCtx
}

// GetStartTime get parents Span Context
func (s metaSpan) GetStartTime() time.Time {
	return s.startTime
}

// SetTag implements ddtrace.Span.
func (metaSpan) SetTag(key string, value interface{}) {}

// SetOperationName implements ddtrace.Span.
func (metaSpan) SetOperationName(operationName string) {}

// BaggageItem implements ddtrace.Span.
func (metaSpan) BaggageItem(key string) string { return "" }

// SetBaggageItem implements ddtrace.Span.
func (metaSpan) SetBaggageItem(key, val string) {}

// Finish implements ddtrace.Span.
func (metaSpan) Finish(opts ...ddtrace.FinishOption) {}

// Tracer implements ddtrace.Span.
func (metaSpan) Tracer() ddtrace.Tracer { return NoopTracer{} }

// Context implements ddtrace.Span.
func (s metaSpan) Context() ddtrace.SpanContext { return s.spanContext }

var _ ddtrace.SpanContext = (*metaSpanContext)(nil)

// metaSpanContext is an implementation of ddtrace.SpanContext that is a no-op.
type metaSpanContext struct {
	spanID  uint64
	traceID uint64
}

// SpanID implements ddtrace.SpanContext.
func (c metaSpanContext) SpanID() uint64 { return c.spanID }

// TraceID implements ddtrace.SpanContext.
func (c metaSpanContext) TraceID() uint64 { return c.traceID }

// ForeachBaggageItem implements ddtrace.SpanContext.
func (metaSpanContext) ForeachBaggageItem(handler func(k, v string) bool) {}
