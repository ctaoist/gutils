package log

import (
	"io"

	Log "github.com/sirupsen/logrus"
)

func init() {
	Log.SetFormatter(&Log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// log.SetLevel(log.DebugLevel)
}

func pre(desc string, args ...any) []any {
	a := []any{"[" + desc + "] "}
	return append(a, args...)
}

func Panic(desc string, args ...any) {
	Log.Panic(pre(desc, args...)...)
}

func Panicf(desc, format string, args ...any) {
	Log.Panicf("["+desc+"] "+format, args...)
}

func Fatal(desc string, args ...any) {
	Log.Fatal(pre(desc, args...)...)
}

func Fatalf(desc, format string, args ...any) {
	Log.Fatalf("["+desc+"] "+format, args...)
}

func Error(desc string, args ...any) {
	Log.Error(pre(desc, args...)...)
}

func Errorf(desc, format string, args ...any) {
	Log.Errorf("["+desc+"] "+format, args...)
}

func Warn(desc string, args ...any) {
	Log.Warn(pre(desc, args...)...)
}

func Warnf(desc, format string, args ...any) {
	Log.Warnf("["+desc+"] "+format, args...)
}

func Info(desc string, args ...any) {
	Log.Info(pre(desc, args...)...)
}

func Infof(desc, format string, args ...any) {
	Log.Infof("["+desc+"] "+format, args...)
}

func Debug(desc string, args ...any) {
	Log.Debug(pre(desc, args...)...)
}

func Debugf(desc, format string, args ...any) {
	Log.Debugf("["+desc+"] "+format, args...)
}

func SetLevel(level Log.Level) {
	Log.SetLevel(level)
}

func SetOutput(out io.Writer) {
	Log.SetOutput(out)
}

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Log.Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)
