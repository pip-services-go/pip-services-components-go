package test_log

import (
	"testing"

	"github.com/pip-services-go/pip-services-components-go/log"
)

func newConsoleLoggerFixture() *LoggerFixture {
	logger := log.NewConsoleLogger()
	fixture := NewLoggerFixture(logger)
	return fixture
}

func TestConsoleLogLevel(t *testing.T) {
	fixture := newConsoleLoggerFixture()
	fixture.TestLogLevel(t)
}

func TestConsoleSimpleLogging(t *testing.T) {
	fixture := newConsoleLoggerFixture()
	fixture.TestSimpleLogging(t)
}

func TestConsoleErrorLogging(t *testing.T) {
	fixture := newConsoleLoggerFixture()
	fixture.TestErrorLogging(t)
}
