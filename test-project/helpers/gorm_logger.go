package helpers

import (
	"context"
	"errors"
	"test-project/constants"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GormLoggerOptions struct {
	Logger                    zerolog.Logger
	LogLevel                  gormlogger.LogLevel
	IgnoreRecordNotFoundError bool
	SlowThreshold             time.Duration
}

type GormLogger struct {
	GormLoggerOptions
}

func NewGormLogger(opts GormLoggerOptions) *GormLogger {
	l := &GormLogger{GormLoggerOptions: opts}
	if l.LogLevel == 0 {
		l.LogLevel = gormlogger.Silent
	}
	return l
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l *GormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		l.Logger.Info().Ctx(ctx).Msgf(s, args...)
	}
}

func (l *GormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		l.Logger.Warn().Ctx(ctx).Msgf(s, args...)
	}
}

func (l *GormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		l.Logger.Error().Ctx(ctx).Msgf(s, args...)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := make(map[string]interface{})
	fields["REF_FILE"] = utils.FileWithLineNum()
	fields["QUERY"] = strings.ReplaceAll(sql, `"`, ``)
	fields["DURATION"] = elapsed.Milliseconds()
	
	if rows == -1 {
		fields["ROW_AFFECTED"] = "-"
	} else {
		fields["ROW_AFFECTED"] = rows
	}

	switch {
	case err != nil && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError) && l.LogLevel >= gormlogger.Error:
		fields["ERROR"] = err.Error()
		l.Logger.Error().Ctx(ctx).Any("GORM", fields).Msg(constants.LOG_QUERY)
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		fields["WARNING"] = "SLOW SQL"
		l.Logger.Warn().Ctx(ctx).Any("GORM", fields).Msg(constants.LOG_QUERY)
	case l.LogLevel == gormlogger.Info:
		l.Logger.Info().Ctx(ctx).Any("GORM", fields).Msg(constants.LOG_QUERY)
	}
}
