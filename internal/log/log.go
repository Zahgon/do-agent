package log

import (
	"log"
	"os"
	"regexp"

	kitlog "github.com/go-kit/kit/log"
)

// Level is a log level such a Debug or Error
type Level int

const (
	syslogFlags = log.Llongfile
	normalFlags = log.LUTC | log.Ldate | log.Ltime | log.Llongfile

	// LevelDebug enables debug logging
	LevelDebug Level = iota
	// LevelError enables error logging
	LevelError Level = iota
)

var (
	debuglog = log.New(os.Stdout, "DEBUG: ", normalFlags)
	errlog   = log.New(os.Stderr, "ERROR: ", normalFlags)

	kitlogLogger  = kitlog.NewNopLogger()
	kitlogTrimmer = regexp.MustCompile("level=[a-zA-Z]+")
	kitlogDebug   = kitlog.NewLogfmtLogger(CollectorWriter{debuglog, kitlogTrimmer})

	level = LevelError
)

// CollectorWriter implements io.Writer.
type CollectorWriter struct {
	stdlog       *log.Logger
	trimmerRegex *regexp.Regexp
}

// Write implements io.Writer.
func (w CollectorWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Adjust the calldepth to account for the kitlog frames

// SetLevel sets the log level
func SetLevel(l Level) {
	_ = "STUB: not implemented"

	// InitSyslog initializes logging to syslog
	return
}

func InitSyslog() (err error) { _ = "STUB: not implemented"; return nil }

// Debug prints a debug message. If syslog is enabled then LOG_NOTICE is used
func Debug(msg string, params ...interface{}) { _ = "STUB: not implemented"; return }

// Error prints an error message. If syslog is enabled then LOG_ERR is used
func Error(msg string, params ...interface{}) { _ = "STUB: not implemented"; return }

// Fatal logs Error and exits 1
func Fatal(msg string, params ...interface{}) { _ = "STUB: not implemented"; return }

func GetCollectorLogger() kitlog.Logger { _ = "STUB: not implemented"; return *new(kitlog.Logger) }
