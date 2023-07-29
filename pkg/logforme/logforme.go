package logforme

type Field interface{}
type Fields map[string]Field
type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
	Panic(string, ...any)
	Fatal(string, ...any)
	WithFields(Fields) Logger
}
