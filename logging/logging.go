package logging

import (
	"context"
	"fmt"
	"log"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Logger is interface for looger that DataDog trace supported.
type Logger interface {
	Sync() error
	InfoContext(ctx context.Context, msg string, keysAndValues ...interface{})
	WarnContext(ctx context.Context, msg string, keysAndValues ...interface{})
	ErrorContext(ctx context.Context, msg string, keysAndValues ...interface{})
	Info(keysAndValues ...interface{})
	Infof(format string, keysAndValues ...interface{})
	Error(keysAndValues ...interface{})
	Errorf(format string, keysAndValues ...interface{})
	Fatal(keysAndValues ...interface{})
	Fatalf(format string, keysAndValues ...interface{})
	InfoSpan(msg string, span tracer.Span, keysAndValues ...interface{})
	WarnSpan(msg string, err error, span tracer.Span, keysAndValues ...interface{})
	ErrorSpan(msg string, err error, span tracer.Span, keysAndValues ...interface{})
	FinishSpanWithError(op string, sp tracer.Span, err error, keysAndValues ...interface{})
	FinishSpan(op string, sp tracer.Span, keysAndValues ...interface{})
	WrapError(sname string, fname string, detail string, err error) error // Wrap error with additional info
}

var ddTraceIDKey = "dd.trace_id"

// Logger ...
type logger struct {
	*zap.SugaredLogger
	stats statsd.ClientInterface
}

// Wrap error with additional info and return error.
func (s *logger) WrapError(sname string, fname string, detail string, err error) error {
	return errors.Wrap(err, "["+sname+"."+fname+"] "+detail)
}

// InfoContext Creates info log with context. Appends dd.trace_id if context has a DD span.
func (s *logger) InfoContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	if span, hasSpan := tracer.SpanFromContext(ctx); hasSpan {
		keysAndValues = prependKeyAndValue(keysAndValues, ddTraceIDKey, span.Context().TraceID())
	}
	s.Infow(msg, keysAndValues...)
}

func (s *logger) WarnContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	if span, hasSpan := tracer.SpanFromContext(ctx); hasSpan {
		keysAndValues = prependKeyAndValue(keysAndValues, ddTraceIDKey, span.Context().TraceID())
	}
	s.Warnw(msg, keysAndValues...)
}

func (s *logger) ErrorContext(ctx context.Context, msg string, keysAndValues ...interface{}) {
	if span, hasSpan := tracer.SpanFromContext(ctx); hasSpan {
		keysAndValues = prependKeyAndValue(keysAndValues, ddTraceIDKey, span.Context().TraceID())
	}
	s.Errorw(msg, keysAndValues...)
}

// InfoSpan ...
func (s *logger) InfoSpan(msg string, span tracer.Span, keysAndValues ...interface{}) {
	k := prependKeyAndValue(keysAndValues, ddTraceIDKey, span.Context().TraceID())
	s.Infow(msg, k...)
}

// WarnSpan ...
func (s *logger) WarnSpan(msg string, err error, span tracer.Span, keysAndValues ...interface{}) {
	k := prependKeyAndValue(keysAndValues, ddTraceIDKey, span.Context().TraceID())
	s.Warnf(fmt.Sprintf("%s : %s", msg, err), k...)
}

// ErrorSpan ...
func (s *logger) ErrorSpan(msg string, err error, span tracer.Span, keysAndValues ...interface{}) {
	k := prependKeyAndValue(keysAndValues, ddTraceIDKey, span.Context().TraceID())
	s.Errorw(fmt.Sprintf("%s : %s", msg, err), k...)
}

// FinishSpanWithError ...
func (s *logger) FinishSpanWithError(op string, sp tracer.Span, err error, keysAndValues ...interface{}) {
	s.ErrorSpan("Error "+op, err, sp, keysAndValues...)
	sp.Finish(tracer.WithError(err))
}

// FinishSpan ...
func (s *logger) FinishSpan(op string, sp tracer.Span, keysAndValues ...interface{}) {
	s.InfoSpan(op+" completed successfully.", sp, keysAndValues...)
	if s.stats != nil {
		if err := s.stats.Incr(op+"_count", []string{}, 1); err != nil {
			s.WarnSpan("failed to increment on datadog when "+op, err, sp, keysAndValues...)
		}
	}
	sp.Finish()
}

func prependKeyAndValue(x []interface{}, k interface{}, v interface{}) []interface{} {
	x = append(x, 0)
	x = append(x, 0)
	copy(x[2:], x)
	x[0] = k
	x[1] = v
	return x
}

// NewLogger ...
func NewLogger(output string) Logger {
	return &logger{
		SugaredLogger: NewLoggerWithLevel(output, "info"),
	}
}

// NewLoggerWithStatsd ...
func NewLoggerWithStatsd(output string, clientInterface statsd.ClientInterface) Logger {
	return &logger{
		SugaredLogger: NewLoggerWithLevel(output, "info"),
		stats:         clientInterface,
	}
}

// NewLoggerWithLevel ...
func NewLoggerWithLevel(output string, level string) *zap.SugaredLogger {
	lvl := zap.AtomicLevel{}
	err := lvl.UnmarshalText([]byte((level)))
	if err != nil {
		log.Fatalf("couldn't create logger, err:%s", err)
	}
	cfg := zap.Config{
		Level:       lvl,
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{output},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}
