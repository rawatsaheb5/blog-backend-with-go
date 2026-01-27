package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Logger is the global logger instance
	Logger *zap.Logger
	// Sugar is a convenience wrapper around Logger for easier usage
	Sugar *zap.SugaredLogger
)

// Environment types
const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
	EnvTest        = "test"
)

// Init initializes the logger based on the environment
func Init(env string) error {
	var config zap.Config
	var err error

	// Normalize environment string
	env = strings.ToLower(strings.TrimSpace(env))
	if env == "" {
		env = EnvDevelopment // Default to development
	}

	switch env {
	case EnvProduction:
		config = getProductionConfig()
	case EnvTest:
		config = getTestConfig()
	default: // Development
		config = getDevelopmentConfig()
	}

	// Build the logger
	Logger, err = config.Build()
	if err != nil {
		return err
	}

	// Create sugared logger for convenience
	Sugar = Logger.Sugar()

	// Replace global logger
	zap.ReplaceGlobals(Logger)

	return nil
}

// getDevelopmentConfig returns a development-friendly zap configuration
func getDevelopmentConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    getDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// getProductionConfig returns a production-optimized zap configuration
func getProductionConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    getProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// getTestConfig returns a test-friendly zap configuration
func getTestConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.WarnLevel),
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    getTestEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// getDevelopmentEncoderConfig returns encoder config for development
func getDevelopmentEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// getProductionEncoderConfig returns encoder config for production
func getProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// getTestEncoderConfig returns encoder config for testing
func getTestEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      zapcore.OmitKey,
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// GetLogger returns the global logger instance
func GetLogger() *zap.Logger {
	if Logger == nil {
		// Fallback: initialize with development config if not initialized
		_ = Init(EnvDevelopment)
	}
	return Logger
}

// GetSugar returns the global sugared logger instance
func GetSugar() *zap.SugaredLogger {
	if Sugar == nil {
		// Fallback: initialize with development config if not initialized
		_ = Init(EnvDevelopment)
	}
	return Sugar
}

// Sync flushes any buffered log entries
func Sync() error {
	if Logger != nil {
		return Logger.Sync()
	}
	return nil
}

// GetEnvironment returns the current environment from ENV variable
func GetEnvironment() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("ENVIRONMENT")
	}
	if env == "" {
		return EnvDevelopment
	}
	return strings.ToLower(strings.TrimSpace(env))
}
