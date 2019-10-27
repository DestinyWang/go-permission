package database

import (
	"github.com/DestinyWang/go-permission/util"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAcl(t *testing.T) {
	util.InitTestConfig()
	acl := &Acl{
		Id:          0,
		Code:        "",
		Name:        "name",
		AclModuleId: 0,
		Url:         "url",
		Type:        0,
		Status:      0,
		Seq:         0,
		Remark:      "first acl",
		Operator:    "destinywk",
		//OperateTime: 0,
		OperateIp:   "127.0.0.1",
	}
	err := AddAcl(acl)
	assert.Nil(t, err)
}

func TestFindById(t *testing.T) {
	util.InitTestConfig()
	acl, err := FindById(1)
	assert.Nil(t, err)
	logrus.Info(acl)
}
