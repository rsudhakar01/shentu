package logrus

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/tendermint/tendermint/libs/log"
)

type LogrusLogger struct {
	logger logrus.Logger
}

func NewLogrusLogger() log.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	return &LogrusLogger{
		*logger,
	}
}

// interface assertion
var _ log.Logger = (*LogrusLogger)(nil)

func (l *LogrusLogger) Debug(msg string, keyvals ...interface{}) {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, errors.New("(MISSING)"))
	}
	fields := logrus.Fields{}
	for i := 0; i < len(keyvals); i += 2 {
		fields[fmt.Sprintf("%v", keyvals[i])] = keyvals[i+1]
	}
	l.logger.WithFields(fields).Debug(msg)
}

func (l *LogrusLogger) Info(msg string, keyvals ...interface{}) {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, errors.New("(MISSING)"))
	}
	fields := logrus.Fields{}
	for i := 0; i < len(keyvals); i += 2 {
		fields[fmt.Sprintf("%v", keyvals[i])] = keyvals[i+1]
	}
	l.logger.WithFields(fields).Info(msg)
}

func (l *LogrusLogger) Error(msg string, keyvals ...interface{}) {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, errors.New("(MISSING)"))
	}
	fields := logrus.Fields{}
	for i := 0; i < len(keyvals); i += 2 {
		fields[fmt.Sprintf("%v", keyvals[i])] = keyvals[i+1]
	}
	l.logger.WithFields(fields).Error(msg)
}

func (l *LogrusLogger) With(keyvals ...interface{}) log.Logger {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, errors.New("(MISSING)"))
	}
	fields := logrus.Fields{}
	for i := 0; i < len(keyvals); i += 2 {
		fields[fmt.Sprintf("%v", keyvals[i])] = keyvals[i+1]
	}
	l.logger.WithFields(fields)
	return l
}
