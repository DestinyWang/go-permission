package service

import (
	"github.com/DestinyWang/go-permission/util"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTree(t *testing.T) {
	util.InitTestConfig()
	deptLevelDTOs, err := DepartmentTree()
	assert.Nil(t, err)
	t.Log(deptLevelDTOs)
}

func TestStringsIndex(t *testing.T) {
	t.Log(strings.Index("abc", "def"))
}
