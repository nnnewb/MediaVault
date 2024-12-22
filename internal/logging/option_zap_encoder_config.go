package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// WithZapEncoderConfigMessageKey Set the keys used for each log entry. If any key is empty, that portion
// of the entry is omitted.
func WithZapEncoderConfigMessageKey(MessageKey string) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.MessageKey = MessageKey
	}
}

func WithZapEncoderConfigLevelKey(LevelKey string) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.LevelKey = LevelKey
	}
}

func WithZapEncoderConfigTimeKey(TimeKey string) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.TimeKey = TimeKey
	}
}

func WithZapEncoderConfigNameKey(NameKey string) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.NameKey = NameKey
	}
}

func WithZapEncoderConfigCallerKey(CallerKey string) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.CallerKey = CallerKey
	}
}

func WithZapEncoderConfigFunctionKey(FunctionKey string) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.FunctionKey = FunctionKey
	}
}

func WithZapEncoderConfigStacktraceKey(StacktraceKey string) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.StacktraceKey = StacktraceKey
	}
}

func WithZapEncoderConfigSkipLineEnding(SkipLineEnding bool) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.SkipLineEnding = SkipLineEnding
	}
}

func WithZapEncoderConfigLineEnding(LineEnding string) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.LineEnding = LineEnding
	}
}

// WithZapEncoderConfigEncodeLevel configure the primitive representations of common complex types. For
// example, some users may want all time.Times serialized as floating-point
// seconds since epoch, while others may prefer ISO8601 strings.
func WithZapEncoderConfigEncodeLevel(EncodeLevel zapcore.LevelEncoder) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.EncodeLevel = EncodeLevel
	}
}

func WithZapEncoderConfigEncodeTime(EncodeTime zapcore.TimeEncoder) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.EncodeTime = EncodeTime
	}
}

func WithZapEncoderConfigEncodeDuration(EncodeDuration zapcore.DurationEncoder) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.EncodeDuration = EncodeDuration
	}
}

func WithZapEncoderConfigEncodeCaller(EncodeCaller zapcore.CallerEncoder) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.EncodeCaller = EncodeCaller
	}
}

// WithZapEncoderConfigEncodeName unlike the other primitive type encoders, EncodeName is optional. The
// zero value falls back to FullNameEncoder.
func WithZapEncoderConfigEncodeName(EncodeName zapcore.NameEncoder) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.EncodeName = EncodeName
	}
}

// WithZapEncoderConfigConsoleSeparator configures the field separator used by the console encoder. Defaults
// to tab.
func WithZapEncoderConfigConsoleSeparator(ConsoleSeparator string) Option {
	return func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig) {
		config.EncoderConfig.ConsoleSeparator = ConsoleSeparator
	}
}
