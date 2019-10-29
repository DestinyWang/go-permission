package database

import (
	"fmt"
	"github.com/DestinyWang/go-permission/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetByIdDB(t *testing.T) {
	util.InitTestConfig()
	department, err := GetDeptById(13)
	assert.Nil(t, err)
	fmt.Println(department)
}

func TestGetChildDeptByLevel(t *testing.T) {
	util.InitTestConfig()
	deptList, err := GetChildDeptByLevel("0")
	assert.Nil(t, err)
	t.Log(deptList)
}

func TestUpdateDept(t *testing.T) {
	util.InitTestConfig()
	d := &Department{
		Id: 20,
		Name: "destinywk-sub",
	}
	err := UpdateDept(d)
	assert.Nil(t, err)
}

func TestCountByParentIdAndName(t *testing.T) {
	util.InitTestConfig()
	cnt, err := CountByParentIdAndName(18, "destinywk-sub")
	assert.Nil(t, err)
	t.Log(cnt)
}
