package mysql

import (
	"context"
	"fmt"
	"test/configs"
	"test/helpers"
	hModels "test/helpers/models"
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

type MysqlDatabase interface {
	HealthCheck(ctx context.Context) hModels.DataHealthCheck
	GetDB() *gorm.DB
}

type mysqlDatabase struct {
	Db   *gorm.DB
	Logs zerolog.Logger
}

func (m *mysqlDatabase) GetDB() *gorm.DB {
	return m.Db
}

func InitializeMysqlDatabase(conn *gorm.DB, log zerolog.Logger) MysqlDatabase {
	return &mysqlDatabase{
		Db:   conn,
		Logs: log,
	}
}

func ConnectMysql(log zerolog.Logger) *gorm.DB {
	conf := configs.Cfg.Database.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: helpers.NewGormLogger(helpers.GormLoggerOptions{
			Logger:                    log,
			LogLevel:                  gormlogger.Info,
			IgnoreRecordNotFoundError: true,
			SlowThreshold:             200 * time.Millisecond,
		}),
	})
	if err != nil {
		log.Error().Err(err).Msg("Error open mysql connection")
		panic("failed to connect database")
	}
	return db
}

func (m *mysqlDatabase) HealthCheck(ctx context.Context) hModels.DataHealthCheck {
	res := hModels.DataHealthCheck{
		ServiceName: "MySQL Database",
		StatusCode:  200,
	}
	
	sqlDB, err := m.Db.DB()
	if err != nil {
		m.Logs.Error().Ctx(ctx).Err(err).Msg("Failed to get SQL DB")
		res.StatusCode = 500
		res.AdditionalData = err.Error()
		return res
	}
	
	if err := sqlDB.Ping(); err != nil {
		m.Logs.Error().Ctx(ctx).Err(err).Msg("MySQL ping failed")
		res.StatusCode = 500
		res.AdditionalData = err.Error()
		return res
	}
	
	return res
}
