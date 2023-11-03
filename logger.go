package goll

import (
	"github.com/sirupsen/logrus"
	"log"
)

var Log LoggerInterface

// Logger структура, реализующая LoggerInterface на основе logrus
type Logger struct {
	logger *logrus.Logger
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}
func (l *Logger) Traceln(args ...interface{}) {
	l.logger.Traceln(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}
func (l *Logger) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}
func (l *Logger) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}
func (l *Logger) Warnln(args ...interface{}) {
	l.logger.Warnln(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}
func (l *Logger) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}
func (l *Logger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}

// NewLogger создает новый экземпляр Logger с заданными настройками
func NewLogger(level logrus.Level) LoggerInterface {
	logger := logrus.New()
	logger.SetLevel(level)
	logger.SetFormatter(&CustomJSONFormatter{OrderedFields: []string{"level", "time", "msg", "file", "func"}})
	return &Logger{logger: logger}
}

func InitLogger(level string) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatal(err)
	}
	Log = NewLogger(logLevel)
}
