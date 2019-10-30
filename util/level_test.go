package util

import "testing"

func TestCalLevel(t *testing.T) {
	level := CalLevel("0", 123)
	t.Log(level)
}
