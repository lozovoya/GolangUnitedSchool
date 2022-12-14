package zap

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const CallerSkip = 1

type Logger struct {
	l  *zap.Logger
	sl *zap.SugaredLogger
}

func NewLogger(level, encoding string) (*Logger, error) {
	var err error
	logger := &Logger{}

	zapLevel := zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		return nil, fmt.Errorf("zap logger unmarshal log level: %v", err)
	}

	zapCfg := zap.Config{
		Encoding:         encoding,
		Level:            zapLevel,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "time",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		}}

	logger.l, err = zapCfg.Build(zap.AddCallerSkip(CallerSkip))
	if err != nil {
		return nil, fmt.Errorf("zap logger build: %v", err)
	}

	logger.sl = logger.l.Sugar()

	return logger, nil
}

// Fatal is for Fatal
func (lg *Logger) Fatal(args ...interface{}) {
	lg.sl.Fatal(args...)
}

// Fatalf is for Fatalf
func (lg *Logger) Fatalf(tmpl string, args ...interface{}) {
	lg.sl.Fatalf(tmpl, args...)
}

// Fatalw is for Fatalw
func (lg *Logger) Fatalw(msg string, err interface{}, args ...interface{}) {
	kvs := make([]interface{}, 0, len(args)+2)
	kvs = append(kvs, "error", err)
	kvs = append(kvs, args...)
	lg.sl.Fatalw(msg, kvs...)
}

// Error is for Error
func (lg *Logger) Error(args ...interface{}) {
	lg.sl.Error(args...)
}

// Errorf is for Errorf
func (lg *Logger) Errorf(tmpl string, args ...interface{}) {
	lg.sl.Errorf(tmpl, args...)
}

// Errorw is for Errorw
func (lg *Logger) Errorw(msg string, err interface{}, args ...interface{}) {
	kvs := make([]interface{}, 0, len(args)+2)
	kvs = append(kvs, "error", err)
	kvs = append(kvs, args...)
	lg.sl.Errorw(msg, kvs...)
}

// Warn is for Warn
func (lg *Logger) Warn(args ...interface{}) {
	lg.sl.Warn(args...)
}

// Warnf is for Warnf
func (lg *Logger) Warnf(tmpl string, args ...interface{}) {
	lg.sl.Warnf(tmpl, args...)
}

// Warnw is for Warnw
func (lg *Logger) Warnw(msg string, args ...interface{}) {
	lg.sl.Warnw(msg, args...)
}

// Info is for Info
func (lg *Logger) Info(args ...interface{}) {
	lg.sl.Info(args...)
}

// Infof is for Infof
func (lg *Logger) Infof(tmpl string, args ...interface{}) {
	lg.sl.Infof(tmpl, args...)
}

// Infow is for Infow
func (lg *Logger) Infow(msg string, args ...interface{}) {
	lg.sl.Infow(msg, args...)
}

// Debug is for Debug
func (lg *Logger) Debug(args ...interface{}) {
	lg.sl.Debug(args...)
}

// Debugf is for Debugf
func (lg *Logger) Debugf(tmpl string, args ...interface{}) {
	lg.sl.Debugf(tmpl, args...)
}

// Debugw is for Debugw
func (lg *Logger) Debugw(msg string, args ...interface{}) {
	lg.sl.Debugw(msg, args...)
}

// Sync is for sync
func (lg *Logger) Sync() {
	err := lg.sl.Sync()
	if err != nil {
		log.Println("Fail to sync zap-logger", err)
	}
}

func (lg *Logger) With(arg ...interface{}) {
	lg.sl = lg.sl.With(arg)
}
