package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config Cfg

type Cfg struct {
	Database struct{
        UserName string
        Password string
        Host string
        Port int
        DbName string
        Url string
	}
}

func InitConfig(configPath string) (err error) {
	viper.SetDefault("env", "test") // 设置默认环境
	if err = viper.BindEnv("env", "env"); err != nil { // 绑定系统环境
		logrus.WithError(err).Error("bind env fail")
		return
	}
	env := viper.GetString("env")
	cfgFileName := fmt.Sprintf("conf-%s", env)
	logrus.Infof("config file=[%s]", cfgFileName)
	viper.SetConfigName(cfgFileName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yml")
	if err = viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Error("read config fail")
		return
	}
	if err = viper.Unmarshal(&Config); err != nil {
		logrus.WithError(err).Error("unmarshal Config fail")
		return
	}
	logrus.Infof("config init succ: [%+v]", Config)
	return nil
}
