package logger

type LogFields map[string]any

type Logger interface {
	Info(msg string, fields *LogFields)
	Warn(msg string, fields *LogFields)
	Error(msg string, fields *LogFields)
	Fatal(msg string, fields *LogFields)
	Panic(msg string, fields *LogFields)
}

var GlobalLogger Logger

func Init() {
	GlobalLogger = NewLogrusLogger()
}
