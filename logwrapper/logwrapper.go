package logwrapper

import (
	"io"
	"os"
	"time"

	"github.com/abruneau/pi_custom_metrics/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {

	var writers []io.Writer

	if config.Conf.GetBool("log_to_file") {
		file := config.Conf.GetString("log_file")
		fWriter, _ := rotatelogs.New(
			file+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(file),
			rotatelogs.WithRotationCount(1),
			// rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
			rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		)
		writers = append(writers, fWriter)
	}

	if config.Conf.GetBool("log_to_console") {
		writers = append(writers, os.Stdout)
	}

	ll, err := logrus.ParseLevel(config.GetLogLevel())
	if err != nil {
		ll = logrus.DebugLevel
	}
	var baseLogger = logrus.New()
	baseLogger.SetLevel(ll)

	var standardLogger = &StandardLogger{baseLogger}
	standardLogger.SetOutput(io.MultiWriter(writers...))
	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	invalidArgMessage      = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid value for argument: %s: %v"}
	missingArgMessage      = Event{3, "Missing arg: %s"}
)

// InvalidArg is a standard error message
func (l *StandardLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *StandardLogger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage.message, argumentName)
}
