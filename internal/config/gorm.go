package config

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

func NewDatabase(config *viper.Viper, log *logrus.Logger) (*gorm.DB, error) {
	host := config.GetString("DB_HOST")
	port := config.GetString("DB_PORT")
	username := config.GetString("DB_USERNAME")
	password := config.GetString("DB_PASSWORD")
	database := config.GetString("DB_DATABASE")

	idleConnection := config.GetInt("DB_POOL_IDLE")
	maxConnection := config.GetInt("DB_POOL_MAX")
	maxLifeTimeConnection := config.GetInt("DB_POOL_LIFETIME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable timezone=UTC", host, username, password, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger(log),
	})
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}

	connection, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxIdleConns(maxConnection)
	connection.SetConnMaxIdleTime(time.Second * time.Duration(maxLifeTimeConnection))

	return db, nil
}

type gormLogger struct {
	Log                   *logrus.Logger
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Debug                 bool
}

func newLogger(log *logrus.Logger) *gormLogger {
	return &gormLogger{
		Log:                   log,
		SkipErrRecordNotFound: true,
		Debug:                 true,
		SlowThreshold:         time.Second * 5,
	}
}

func (l *gormLogger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (l *gormLogger) Info(ctx context.Context, s string, args ...any) {
	l.Log.WithContext(ctx).Infof(s, args...)
}

func (l *gormLogger) Warn(ctx context.Context, s string, args ...any) {
	l.Log.WithContext(ctx).Warnf(s, args...)
}

func (l *gormLogger) Error(ctx context.Context, s string, args ...any) {
	l.Log.WithContext(ctx).Errorf(s, args...)
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		l.Log.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.Log.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	if l.Debug {
		l.Log.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
	}
}
