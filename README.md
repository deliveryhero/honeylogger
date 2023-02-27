![Version](https://img.shields.io/badge/version-1.3.3-orange.svg)
[![GolangCI Lint](https://github.com/deliveryhero/honeylogger/actions/workflows/go-lint.yml/badge.svg)](https://github.com/deliveryhero/honeylogger/actions/workflows/go-lint.yml)
[![Golang Tests](https://github.com/deliveryhero/honeylogger/actions/workflows/go-test.yml/badge.svg)](https://github.com/deliveryhero/honeylogger/actions/workflows/go-test.yml)

# Honeylogger

Simple logger with DataDog’s **span** support. Uses
[Zap](https://github.com/uber-go/zap) under the hood.

---

## Installation

Now you can add this package via;

```bash
go get github.com/deliveryhero/honeylogger
```

---

## GolangCI Linter Settings

You can add these lines to your existing linter yaml if you use `wrapcheck` linter:

```yaml
linters-settings:
  wrapcheck:
    ignoreSigs:
      - .Errorf(
      - errors.New(
      - errors.Unwrap(
      - .Wrap(
      - .Wrapf(
      - .WithMessage(
      - .WithMessagef(
      - .WithStack(
      - .WrapError(
```

---

## Usage

```go
logger := logging.NewLogger("stderr")
logger.Fatal(errors.Wrap(err, "invalid server port"))
```

## With context

This adds **span** information if context has **DataDog Span** object.

```go
logger := logging.NewLogger("stderr")
span, spanCtx := tracer.StartSpanFromContext(...)
defer func() {
	span.Finish(tracer.WithError(err))
}()
    
logger.InfoContext(spanCtx, fmt.Sprintf("[outboxService.ticked] non published count: %v", count))
```

## Direct use with DataDog Span

```go
logger := logging.NewLogger("stderr")
   
span, spanCtx := tracer.StartSpanFromContext(...)
defer func() {
	span.Finish(tracer.WithError(err))
}()
    
logger.InfoSpan( fmt.Sprintf("[outboxService.ticked] non published count: %v",count), span)
```

---

## Rake Tasks

```bash
rake -T

rake bump[revision]     # bump version, default is: patch
rake default            # default task
rake doc[port]          # run doc server
rake mockery            # run mockery
rake publish[revision]  # publish new version of the library, default is: patch
rake test               # run tests
```

---

## Tests

To run tests, use `rake test` or;

```bash
go test -p 1 -v -race ./...
```

### Mock Usage

```go
logger := &mocks.Logger{}
logger.On("Info", mock.Anything).Return()
logger.On("InfoSpan", mock.Anything, mock.Anything, mock.Anything).Return()
logger.On("InfoContext", mock.Anything, mock.Anything, mock.Anything).Return()
```

---

## Godoc Server

Make sure you have already installed `godoc` unless:

```bash
go install golang.org/x/tools/cmd/godoc@latest
godoc -http=:9009
```

or, use rake tasks:

- `rake doc` uses default port which is `9009`
- `rake doc[8008]` uses given `8008` as port number

then;

```bash
# for default port
open http://localhost:9009/pkg/github.com/deliveryhero/honeylogger/logging/
```

---

## Publishing New Release

Prerequisites

- You need to be in `main` branch
- You need to be ready to bump to a new version

Use `rake publish[revision]` task to bump new version and push newly created
tag and updated code to remote and verify go package. (all in one!)

- `rake publish`: `0.0.0` -> `0.0.1`, default revision is `patch`
- `rake publish[minor]`: `0.0.0` -> `0.1.0`
- `rake publish[major]`: `0.0.0` -> `1.0.0`

---

## Contributor(s)

* [Hakan Kutluay](https://github.com/hakankutluay) - Creator, maintainer
* [Erhan Akpınar](https://github.com/erhanakp) - Contributor
* [Uğur "vigo" Özyılmazel](https://github.com/vigo) - Contributor

---


## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/deliveryhero/honeylogger/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'add some functionality'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [code of conduct][coc].


[coc]: https://github.com/deliveryhero/honeylogger/blob/main/CODE_OF_CONDUCT.md