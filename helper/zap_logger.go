package helper

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var SugarObj *SugaredLogger
var LoggerObj *zap.Logger

// SugaredLogger is a wrapper around zap's SugaredLogger
type SugaredLogger struct {
	sugarObj *zap.SugaredLogger
}

// init initializes the logger during the init phase
func init() {
	// Log rotation configuration using lumberjack
	lumberjackLogrotate := &lumberjack.Logger{
		Filename:   "logs/access.log",
		MaxSize:    10,  // Max megabytes before log is rotated
		MaxBackups: 90,  // Max number of old log files to keep
		MaxAge:     180, // Max number of days to retain log files
		Compress:   true,
	}
	writerSyncer := zapcore.AddSync(lumberjackLogrotate)

	// Set up the encoder configuration for structured JSON logging
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Set log level based on environment (defaults to InfoLevel)
	level := zapcore.InfoLevel
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel == "debug" {
		level = zapcore.DebugLevel
	}

	// Define the log level and core configuration for the logger
	atom := zap.NewAtomicLevel()
	atom.SetLevel(level)

	// Create the logger object with the core, adding caller information for debugging
	// AddCallerSkip(1) skips over the first frame to avoid capturing the file where logger is created
	LoggerObj = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(writerSyncer),
		atom,
	), zap.AddCaller(), zap.AddCallerSkip(1)) // Skip 1 stack frame

	// Create the sugared logger for easier logging
	SugarObj = &SugaredLogger{
		sugarObj: LoggerObj.Sugar(),
	}

	// Ensure logs are flushed properly on shutdown
	defer func() {
		if err := LoggerObj.Sync(); err != nil {
			log.Printf("Failed to flush logs: %v", err)
		}
	}()
}

// Debug logs a message at debug level
func (s *SugaredLogger) Debug(args ...interface{}) {
	s.sugarObj.Debug(args...)
}

// Info logs a message at info level
func (s *SugaredLogger) Info(args ...interface{}) {
	s.sugarObj.Info(args...)
}

// Warn logs a message at warn level
func (s *SugaredLogger) Warn(args ...interface{}) {
	s.sugarObj.Warn(args...)
}

// Error logs a message at error level
func (s *SugaredLogger) Error(args ...interface{}) {
	s.sugarObj.Error(args...)
}
