package logforme

import "github.com/sirupsen/logrus"

type logrusEntry struct {
	entry *logrus.Entry
}
type logrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusEntry(entry *logrus.Entry) (Logger, error) {
	return &logrusEntry{
		entry: entry,
	}, nil
}

func NewLogrusLogger(logger *logrus.Logger) (Logger, error) {
	return &logrusLogger{
		logger: logger,
	}, nil
}

func (l *logrusLogger) Debug(tpl string, args ...any) {
	l.logger.Debugf(tpl, args...)
}
func (l *logrusLogger) Info(tpl string, args ...any) {
	l.logger.Infof(tpl, args...)
}
func (l *logrusLogger) Warn(tpl string, args ...any) {
	l.logger.Warnf(tpl, args...)
}

func (l *logrusLogger) Error(tpl string, args ...any) {
	l.logger.Errorf(tpl, args...)
}

func (l *logrusLogger) Panic(tpl string, args ...any) {
	l.logger.Panicf(tpl, args...)
}

func (l *logrusLogger) Fatal(tpl string, args ...any) {
	l.logger.Fatalf(tpl, args...)
}

func (l *logrusLogger) WithFields(fields Fields) Logger {
	return &logrusEntry{
		entry: l.logger.WithFields(toLogrusFields(fields)),
	}
}

func (l *logrusEntry) Debug(tpl string, args ...any) {
	l.entry.Debugf(tpl, args...)
}
func (l *logrusEntry) Info(tpl string, args ...any) {
	l.entry.Infof(tpl, args...)
}
func (l *logrusEntry) Warn(tpl string, args ...any) {
	l.entry.Warnf(tpl, args...)
}
func (l *logrusEntry) Error(tpl string, args ...any) {
	l.entry.Errorf(tpl, args...)
}
func (l *logrusEntry) Panic(tpl string, args ...any) {
	l.entry.Panicf(tpl, args...)
}
func (l *logrusEntry) Fatal(tpl string, args ...any) {
	l.entry.Fatalf(tpl, args...)
}
func (l *logrusEntry) WithFields(fields Fields) Logger {
	return &logrusEntry{
		entry: l.entry.WithFields(toLogrusFields(fields)),
	}
}

func toLogrusFields(f Fields) logrus.Fields {
	lf := logrus.Fields{}
	for index, val := range f {
		lf[index] = val
	}
	return lf
}
