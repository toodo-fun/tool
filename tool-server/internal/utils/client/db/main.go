package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"path/filepath"
	"time"
	"tool-server/internal/utils/path"
)

var (
	client *gorm.DB
)

func NewDbClient(source string, maxOpenConn int, maxIdleConn int, connMaxLifetime int) *gorm.DB {
	if client == nil {
		var err error
		if !path.IsExist(source) {
			p := filepath.Dir(source)
			err = path.MkdirAll(p)
			if err != nil {
				logrus.Panic(err)
			}
		}
		client, err = gorm.Open(sqlite.Open(source), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		db, _ := client.DB()
		db.SetMaxOpenConns(maxOpenConn)
		db.SetMaxIdleConns(maxIdleConn)
		db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
		if err != nil {
			logrus.Errorf("DB client create failed: %+v", err)
			panic(err)
		} else {
			logrus.Info("DB client create succeed")
		}
	}
	return client
}
