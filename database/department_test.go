package database

import (
	"fmt"
	"github.com/DestinyWang/go-permission/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetByIdDB(t *testing.T) {
	util.InitTestConfig()
	department, err := GetByIdDB(13)
	assert.Nil(t, err)
	fmt.Println(department)
}
