package main

import (
	"github.com/DestinyWang/go-permission/util"
	"runtime"
)

// 初始化
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 解析命令行参数
func initArgs() {
}

// 加载配置项
func initConfig() *util.Config {
	config := new(util.Config)
	return config
}

func main() {
	var err error
	// 初始化命令行参数
	initArgs()
	// 初始化线程
	initEnv()
	initConfig()
}
