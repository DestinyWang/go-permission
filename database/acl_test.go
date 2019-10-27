package database

import (
	"github.com/DestinyWang/go-permission/util"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"testing"
)

var validate *validator.Validate

func TestAddAcl(t *testing.T) {
	util.InitTestConfig()
	var err error
	acl := &Acl{
		//Id:          0,
		//Code:        "",
		//Name:        "name",
		AclModuleId: 0,
		Url:         "url",
		Type:        0,
		Status:      3,
		Seq:         0,
		Remark:      "first acl",
		Operator:    "destinywk",
		OperateTime: util.CurrMillSecond(),
		OperateIp:   "127.0.0.1",
	}
	logrus.Infof("acl=[%+v]", acl)
	validate = validator.New()
	//validate.RegisterStructValidation()
	err = validate.Struct(acl)
	assert.Nil(t, err)
	//err = AddAcl(acl)
	//assert.Nil(t, err)
}

func TestFindById(t *testing.T) {
	util.InitTestConfig()
	acl, err := FindById(1)
	assert.Nil(t, err)
	logrus.Info(acl)
}
