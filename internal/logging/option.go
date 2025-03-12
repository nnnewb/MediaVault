package logging

import (
	"go.uber.org/zap"
)

type Option func(config *zap.Config, lumberjackLogger *LumberjackLoggerConfig)
