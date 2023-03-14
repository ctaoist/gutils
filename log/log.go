package log

import (
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

func Fatal(desc string, args ...any) {
	a := []any{desc}
	a = append(a, args...)
	Log.Fatal(a...)
}

func Error(desc string, args ...interface{}) {
	a := []any{desc}
	a = append(a, args...)
	Log.Error(a...)
}

func Warn(desc string, args ...interface{}) {
	a := []any{desc}
	a = append(a, args...)
	Log.Warn(a...)
}

func Info(desc string, args ...interface{}) {
	a := []any{desc}
	a = append(a, args...)
	Log.Info(a...)
}

func Debug(desc string, args ...interface{}) {
	a := []any{desc}
	a = append(a, args...)
	Log.Debug(a...)
}

func SetLevel(level Log.Level) {
	Log.SetLevel(level)
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
