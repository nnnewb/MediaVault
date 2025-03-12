package logging

import (
	"net/url"
	"sync/atomic"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger atomic.Pointer[zap.Logger]
var sugar atomic.Pointer[zap.SugaredLogger]

func init() {
	zapLogger, err := zap.NewProduction(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	if err != nil {
		panic(err)
	}
	logger.Store(zapLogger)
	sugar.Store(zapLogger.Sugar())
}

// Setup 日志初始化。
func Setup(options ...Option) error {
	config := zap.NewProductionConfig()
	config.Level.SetLevel(zapcore.InfoLevel)
	config.Sampling = nil
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.OutputPaths = []string{"stderr"}
	lumberjackLoggerConfig := &LumberjackLoggerConfig{
		MaxSize:  100 * 1024 * 1024, // 单一日志文件大小
		MaxAge:   180,               // 等保二级以上均要求日志保存 180 天以上
		Compress: true,              // 降低日志存储空间使用
	}

	for _, option := range options {
		option(&config, lumberjackLoggerConfig)
	}

	err := zap.RegisterSink("lumberjack", func(u *url.URL) (zap.Sink, error) {
		return NewLumberjackSink(u.Opaque, lumberjackLoggerConfig), nil
	})
	if err != nil {
		return errors.WithStack(err)
	}

	// 初始化日志核心
	zapLogger, err := config.Build(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	if err != nil {
		return errors.WithStack(err)
	}

	// 存储日志对象
	logger.Store(zapLogger)
	sugar.Store(zapLogger.Sugar())
	return nil
}

func GetLogger() *zap.Logger {
	return logger.Load()
}

func GetSugar() *zap.SugaredLogger {
	return sugar.Load()
}
