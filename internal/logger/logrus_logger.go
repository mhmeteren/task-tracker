package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	*logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})

	return &LogrusLogger{logger}
}

func (l *LogrusLogger) Info(msg string, fields *LogFields) {
	if fields != nil {
		l.Logger.WithFields(logrus.Fields(*fields)).Info(msg)
	} else {
		l.Logger.Info(msg)
	}
}

func (l *LogrusLogger) Warn(msg string, fields *LogFields) {
	if fields != nil {
		l.Logger.WithFields(logrus.Fields(*fields)).Warn(msg)
	} else {
		l.Logger.Warn(msg)
	}
}

func (l *LogrusLogger) Error(msg string, fields *LogFields) {
	if fields != nil {
		l.Logger.WithFields(logrus.Fields(*fields)).Error(msg)
	} else {
		l.Logger.Error(msg)
	}
}

func (l *LogrusLogger) Fatal(msg string, fields *LogFields) {
	if fields != nil {
		l.Logger.WithFields(logrus.Fields(*fields)).Fatal(msg)
	} else {
		l.Logger.Fatal(msg)
	}
}

func (l *LogrusLogger) Panic(msg string, fields *LogFields) {
	if fields != nil {
		l.Logger.WithFields(logrus.Fields(*fields)).Panic(msg)
	} else {
		l.Logger.Panic(msg)
	}
}
