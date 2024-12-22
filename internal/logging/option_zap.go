package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// WithZapOutputPath sets the output path for the logger.
// empty string will be ignored.
func WithZapOutputPath(path string) Option {
	return func(config *zap.Config, _ *LumberjackLoggerConfig) {
		switch path {
		case "":
			return
		case "stdout", "stderr":
			config.OutputPaths = append(config.OutputPaths, path)
		default:
			// 所有其他路径都视为 lumberjack 封装的轮换日志
			config.OutputPaths = append(config.OutputPaths, "lumberjack:"+path)
		}
	}
}

// WithZapLevel is the minimum enabled logging level. Note that this is a dynamic
// level, so calling Config.Level.SetLevel will atomically change the log
// level of all loggers descended from this config.
func WithZapLevel(lv zapcore.Level) Option {
	return func(config *zap.Config, _ *LumberjackLoggerConfig) {
		config.Level.SetLevel(lv)
	}
}

// WithZapDevelopment puts the logger in development mode, which changes the
// behavior of DPanicLevel and takes stacktraces more liberally.
func WithZapDevelopment(development bool) Option {
	return func(config *zap.Config, _ *LumberjackLoggerConfig) {
		config.Development = development
	}
}

// WithZapDisableCaller stops annotating logs with the calling function's file
// name and line number. By default, all logs are annotated.
func WithZapDisableCaller(disableCaller bool) Option {
	return func(config *zap.Config, _ *LumberjackLoggerConfig) {
		config.DisableCaller = disableCaller
	}
}

// WithZapDisableStacktrace completely disables automatic stacktrace capturing. By
// default, stacktraces are captured for WarnLevel and above logs in
// development and ErrorLevel and above in production.
func WithZapDisableStacktrace(disableStacktrace bool) Option {
	return func(config *zap.Config, _ *LumberjackLoggerConfig) {
		config.DisableStacktrace = disableStacktrace
	}
}

// WithZapSampling sets a sampling policy. A nil SamplingConfig disables sampling.
func WithZapSampling(samplingConfig *zap.SamplingConfig) Option {
	return func(config *zap.Config, _ *LumberjackLoggerConfig) {
		config.Sampling = samplingConfig
	}
}

// WithZapEncoding sets the logger's encoding. Valid values are "json" and
// "console", as well as any third-party encodings registered via
// RegisterEncoder.
func WithZapEncoding(encoding string) Option {
	return func(config *zap.Config, _ *LumberjackLoggerConfig) {
		config.Encoding = encoding
	}
}

// WithZapInitialFields is a collection of fields to add to the root logger.
func WithZapInitialFields(initialFields map[string]interface{}) Option {
	return func(config *zap.Config, _ *LumberjackLoggerConfig) {
		config.InitialFields = initialFields
	}
}
