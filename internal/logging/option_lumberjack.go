package logging

import (
	"go.uber.org/zap"
)

// WithLumberjackMaxSize set the maximum size in megabytes of the log file before it gets
// rotated. It defaults to 100 megabytes.
func WithLumberjackMaxSize(maxSize int) Option {
	return func(_ *zap.Config, lumberjackSinkFactory *LumberjackLoggerConfig) {
		lumberjackSinkFactory.MaxSize = maxSize
	}
}

// WithLumberjackMaxAge set the maximum number of days to retain old log files based on the
// timestamp encoded in their filename.  Note that a day is defined as 24
// hours and may not exactly correspond to calendar days due to daylight
// savings, leap seconds, etc. The default is not to remove old log files
// based on age.
func WithLumberjackMaxAge(maxAge int) Option {
	return func(_ *zap.Config, lumberjackSinkFactory *LumberjackLoggerConfig) {
		lumberjackSinkFactory.MaxAge = maxAge
	}
}

// WithLumberjackMaxBackups set the maximum number of old log files to retain.  The default
// is to retain all old log files (though MaxAge may still cause them to get
// deleted.)
func WithLumberjackMaxBackups(maxBackups int) Option {
	return func(_ *zap.Config, lumberjackSinkFactory *LumberjackLoggerConfig) {
		lumberjackSinkFactory.MaxBackups = maxBackups
	}
}

// WithLumberjackLocalTime determines if the time used for formatting the timestamps in
// backup files is the computer's local time.  The default is to use UTC
// time.
func WithLumberjackLocalTime(localTime bool) Option {
	return func(_ *zap.Config, lumberjackSinkFactory *LumberjackLoggerConfig) {
		lumberjackSinkFactory.LocalTime = localTime
	}
}

// WithLumberjackCompress determines if the rotated log files should be compressed
// using gzip. The default is not to perform compression.
func WithLumberjackCompress(compress bool) Option {
	return func(_ *zap.Config, lumberjackSinkFactory *LumberjackLoggerConfig) {
		lumberjackSinkFactory.Compress = compress
	}
}
