package logc

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// GormLogger .
type GormLogger struct {
	log                   Logger
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Silent                bool
}

// LogMode .
func (l *GormLogger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

// Info .
func (l *GormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	newArgs := append([]interface{}{}, s)
	newArgs = append(newArgs, args...)
	l.log.Info(newArgs...)
}

// Warn .
func (l *GormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	newArgs := append([]interface{}{}, s)
	newArgs = append(newArgs, args...)
	l.log.Warn(newArgs...)
}

// Error .
func (l *GormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	newArgs := append([]interface{}{}, s)
	newArgs = append(newArgs, args...)
	l.log.Error(newArgs...)
}

// Trace .
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if !l.Silent {
		elapsed := time.Since(begin)
		sql, _ := fc()
		fields := logrus.Fields{}
		if l.SourceField != "" {
			fields[l.SourceField] = utils.FileWithLineNum()
		}
		if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
			l.log.WithError(err).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
			return
		}

		if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
			l.log.WithFields(fields).Debugf("%s [%s]", sql, elapsed)
			return
		}

		l.log.WithFields(fields).Tracef("%s [%s]", sql, elapsed)
	}
}

// NewGormLogger .
func NewGormLogger(logger Logger) *GormLogger {
	return &GormLogger{
		log:           logger,
		SlowThreshold: time.Second * time.Duration(1),
		SourceField:   "src",
		Silent:        false,
	}
}
