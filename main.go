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

// 解析命令行参数
func initArgs() {
}

func main() {
	var err error
	// 初始化命令行参数
	initArgs()
	// 初始化线程
	initEnv()
	if err = util.InitDatabase(); err != nil {
		logrus.WithError(err).Errorf("init database fail")
	}
	router := util.InitRouter()
	if err = router.Run(); err != nil {
		logrus.WithError(err).Error("web router init fail")
	}
}
