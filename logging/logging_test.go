package logging_test

import (
	"reflect"
	"testing"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/deliveryhero/sc-honeylogger/logging"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func TestNewLogger(t *testing.T) {
	type args struct {
		output string
	}
	var tests []struct {
		name string
		args args
		want *zap.SugaredLogger
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := logging.NewInfoLogger(tt.args.output); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want  %v", got, tt.want)
			}
		})
	}
}

func TestNewLoggerWithLevel(t *testing.T) {
	type args struct {
		output string
		level  string
	}
	var tests []struct {
		name string
		args args
		want *zap.SugaredLogger
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := logging.NewLoggerWithLevel(tt.args.output, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoggerWithLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleNewInfoLogger() {
	logger := logging.NewInfoLogger("stderr")
	logger.Info("hello")

	err := errors.New("demo error")
	logger.Fatal(errors.Wrap(err, "invalid server port"))
}

func ExampleNewLoggerWithLevel() {
	logger := logging.NewLoggerWithLevel("stderr", "info")
	logger.Info("hello")
}

func ExampleNewLoggerWithInfoStatsd() {
	var DataDogClient statsd.ClientInterface

	logger := logging.NewLoggerWithInfoStatsd("stderr", DataDogClient)
	logger.Info("hello")
}
