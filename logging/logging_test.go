package logging_test

import (
	"reflect"
	"testing"

	"github.com/deliveryhero/sc-honeylogger/logging"
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
			if got := logging.NewLogger(tt.args.output); !reflect.DeepEqual(got, tt.want) {
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
