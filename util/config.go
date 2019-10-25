package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct{
		Url string
	}
	Redis struct{
		Ips string
	}
}

var Cfg Config

func InitConfig() (err error) {
	if err = viper.BindEnv("env", "env"); err != nil {
		logrus.WithError(err).Error("viper bind env fail")
		return
	}
	viper.SetDefault("env", "test")
	env := viper.GetString("env")
	confFile := fmt.Sprintf("config-%s", env)
	logrus.Infof("config file name [%s]", confFile)
	viper.SetConfigName(confFile)
	viper.AddConfigPath("../conf") // 设置路径 ./conf
	viper.SetConfigType("yaml") // 设置后缀
	if err = viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Error("viper read config fail")
		return
	}
	if err = viper.Unmarshal(&Cfg); err != nil {
		logrus.WithError(err).Error("viper unmarshal fail")
		return
	}
	logrus.Infof("config=[%+v]", Cfg)
	return nil
}
