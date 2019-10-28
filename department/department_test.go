package department

import (
	"github.com/DestinyWang/go-permission/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree(t *testing.T) {
	util.InitTestConfig()
	deptLevelDTOs, err := Tree()
	assert.Nil(t, err)
	t.Log(deptLevelDTOs)
}
