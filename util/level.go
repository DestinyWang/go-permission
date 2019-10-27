package util

import (
	"strconv"
	"strings"
)

const (
	LevelSeparator = "/"
	RootLevel      = "0"
)

func CalLevel(parentLevel string, parentId int64) string {
	if parentLevel == "" {
		return RootLevel
	} else {
		return strings.Join([]string{parentLevel, strconv.FormatInt(parentId, 10)}, LevelSeparator)
	}
}
