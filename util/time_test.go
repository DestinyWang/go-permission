package util

import (
	"fmt"
	"testing"
	"time"
)

func TestCurrMillSecond(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Unix())
}
