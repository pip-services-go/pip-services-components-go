package test_log

import (
	"testing"

	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/log"
)

func newCompositeLoggerFixture() *LoggerFixture {
	logger := log.NewCompositeLogger()

	refs := cref.NewReferencesFromTuples(
		cref.NewDescriptor("pip-services", "logger", "console", "default", "1.0"), log.NewConsoleLogger(),
		// cref.NewDescriptor("pip-services", "logger", "composite", "default", "1.0"), logger,
	)
	logger.SetReferences(refs)

	fixture := NewLoggerFixture(logger)
	return fixture
}

func TestCompositeLogLevel(t *testing.T) {
	fixture := newCompositeLoggerFixture()
	fixture.TestLogLevel(t)
}

func TestCompositeSimpleLogging(t *testing.T) {
	fixture := newCompositeLoggerFixture()
	fixture.TestSimpleLogging(t)
}

func TestCompositeErrorLogging(t *testing.T) {
	fixture := newCompositeLoggerFixture()
	fixture.TestErrorLogging(t)
}
