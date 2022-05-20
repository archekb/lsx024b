package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

var std Logger

func init() {
	if std.Logger == nil {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.00000")
		config.EncoderConfig.ConsoleSeparator = " | "
		logger, _ := config.Build()
		std.Logger = logger
	}
}

func ProductionMode() {
	config := zap.NewProductionConfig()
	config.Encoding = "console"
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	config.EncoderConfig.ConsoleSeparator = " | "
	config.DisableCaller = true
	config.DisableStacktrace = true
	logger, _ := config.Build()
	std.Logger = logger
}

func StandartNamed(name string) *Logger {
	return &Logger{Logger: std.Named(name)}
}

// Debug logs a message at level Debug on the standard logger.
func (l *Logger) Debug(args ...interface{}) {
	l.Sugar().Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func (l *Logger) Print(args ...interface{}) {
	l.Sugar().Info(args...)
}

// Info logs a message at level Info on the standard logger.
func (l *Logger) Info(args ...interface{}) {
	l.Sugar().Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l *Logger) Warn(args ...interface{}) {
	l.Sugar().Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func (l *Logger) Warning(args ...interface{}) {
	l.Sugar().Warn(args...)
}

// Error logs a message at level Error on the standard logger.
func (l *Logger) Error(args ...interface{}) {
	l.Sugar().Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func (l *Logger) Panic(args ...interface{}) {
	l.Sugar().Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l *Logger) Fatal(args ...interface{}) {
	l.Sugar().Fatal(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Sugar().Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func (l *Logger) Printf(format string, args ...interface{}) {
	l.Sugar().Infof(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Sugar().Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Sugar().Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.Sugar().Warnf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Sugar().Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.Sugar().Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Sugar().Fatalf(format, args...)
}

// Debugln logs a message at level Debug on the standard logger.
func (l *Logger) Debugln(args ...interface{}) {
	l.Sugar().Debug(args...)
}

// Println logs a message at level Info on the standard logger.
func (l *Logger) Println(args ...interface{}) {
	l.Sugar().Info(args...)
}

// Infoln logs a message at level Info on the standard logger.
func (l *Logger) Infoln(args ...interface{}) {
	l.Sugar().Info(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func (l *Logger) Warnln(args ...interface{}) {
	l.Sugar().Warn(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func (l *Logger) Warningln(args ...interface{}) {
	l.Sugar().Warn(args...)
}

// Errorln logs a message at level Error on the standard logger.
func (l *Logger) Errorln(args ...interface{}) {
	l.Sugar().Error(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func (l *Logger) Panicln(args ...interface{}) {
	l.Sugar().Panic(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l *Logger) Fatalln(args ...interface{}) {
	l.Sugar().Fatal(args...)
}
