package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InitDB() (err error) {
	Db, err := gorm.Open("mysql", Cfg.Database.Url)
	if err != nil {
		logrus.WithError(err).Error("database connection fail")
		return err
	}
	defer Db.Close()
	return nil
}
