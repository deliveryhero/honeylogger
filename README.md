![Version](https://img.shields.io/badge/version-0.0.0-orange.svg)
[![GolangCI Lint](https://github.com/deliveryhero/sc-honeylogger/actions/workflows/go-lint.yml/badge.svg)](https://github.com/deliveryhero/sc-honeylogger/actions/workflows/go-lint.yml)
[![Golang Tests](https://github.com/deliveryhero/sc-honeylogger/actions/workflows/go-test.yml/badge.svg)](https://github.com/deliveryhero/sc-honeylogger/actions/workflows/go-test.yml)
# Social Commerce Honeylogger

Simple logger with DataDog span support. Uses [Zap](https://github.com/uber-go/zap) under the hood.

**Sample usage:**

```go
logger := logging.NewLogger("stderr")
logger.Fatal(errors.Wrap(err, "invalid server port"))
```

**With context (adds span info if context has DataDog span):**

```go
logger := logging.NewLogger("stderr")
   
span, spanCtx := tracer.StartSpanFromContext(...)

defer func() {
	span.Finish(tracer.WithError(err))
}()
    
logger.InfoContext(spanCtx, fmt.Sprintf("[outboxService.ticked] non published count: %v", count))
```

**Direct use with DataDog span:**

```go
logger := logging.NewLogger("stderr")
   
span, spanCtx := tracer.StartSpanFromContext(...)

defer func() {
	span.Finish(tracer.WithError(err))
}()
    
logger.InfoSpan( fmt.Sprintf("[outboxService.ticked] non published count: %v",count), span)
```
## Mock Usage



```go
logger := &mocks.Logger{}
logger.On("Info", mock.Anything).Return()
logger.On("InfoSpan", mock.Anything, mock.Anything, mock.Anything).Return()
logger.On("InfoContext", mock.Anything, mock.Anything, mock.Anything).Return()

```