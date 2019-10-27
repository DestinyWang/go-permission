package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InitDatabase() (err error) {
	url := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8", Config.Database.UserName, Config.Database.Password, Config.Database.Host, Config.Database.Port, Config.Database.DbName)
	logrus.Infof("url of database: [%s]", url)
	Db, err = gorm.Open("mysql", url)
	if err != nil {
		logrus.WithError(err).Error("database connection fail")
		return err
	}
	Db.SingularTable(true)
	//defer Db.Close()
	return nil
}
