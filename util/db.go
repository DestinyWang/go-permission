package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func InitDB() {
	Db, err := gorm.Open("mysql", "root:123456@/auth?charset=utf8")
	if err != nil {
	
	}
	defer Db.Close()
}
