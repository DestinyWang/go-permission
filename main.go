package main

import (
	"github.com/DestinyWang/go-permission/util"
	"github.com/sirupsen/logrus"
	"runtime"
)

// 初始化
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var err error
	// 初始化线程
	if err = util.InitConfig("conf"); err != nil {
		logrus.WithError(err).Error("init config fail")
		return
	}
	initEnv()
	if err = util.InitDatabase(); err != nil {
		logrus.WithError(err).Errorf("init database fail")
		return
	}
	router := InitRouter()
	if err = router.Run(); err != nil {
		logrus.WithError(err).Error("web router init fail")
		return
	}
}
