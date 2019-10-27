package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitConfig(t *testing.T) {
	err := InitConfig("../conf")
	assert.Nil(t, err)
}
