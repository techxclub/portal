package utils

import (
	"strconv"

	"github.com/satori/go.uuid"
)

func ParseInt64WithDefault(str string, def int64) int64 {
	parsedValue, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return def
	}

	return parsedValue
}

func GetRandomUUID() string {
	return uuid.NewV4().String()
}
