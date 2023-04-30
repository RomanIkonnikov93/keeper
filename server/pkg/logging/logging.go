// Package logging using the logrus library.
package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

// writerHook struct.
type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Fire for writerHook struct.
func (hook *writerHook) Fire(entry *logrus.Entry) error {

	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}

	return err
}

// Levels for writerHook struct.
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// Logger for GetLogger.
type Logger struct {
	*logrus.Entry
}

// GetLogger creates a new logger.
func GetLogger() *Logger {
	e := Log()
	return &Logger{
		e,
	}
}

// Log logger implementation.
func Log() *logrus.Entry {

	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0744)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0744)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	e := logrus.NewEntry(l)

	return e
}
