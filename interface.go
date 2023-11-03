package goll

// LoggerInterface определяет методы, которые должен иметь логгер
type LoggerInterface interface {
	Tracef(format string, args ...interface{})
	Traceln(args ...interface{})

	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})

	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
}
