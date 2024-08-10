package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"
	"github.com/techx/portal/errors"
)

func ParseInt64WithDefault(str string, def int64) int64 {
	parsedValue, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return def
	}

	return parsedValue
}

func GenerateRandomNumber(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("invalid length")
	}

	maxNum := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil)
	randomNumber, err := rand.Int(rand.Reader, maxNum)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate random number")
		return "", err
	}

	format := "%0" + strconv.Itoa(length) + "d"
	return fmt.Sprintf(format, randomNumber.Int64()), nil
}

func GetRandomUUID() string {
	return uuid.NewV4().String()
}

func UpdatedToZeroValue[T comparable](oldValue, newValue T) bool {
	var zero T
	if oldValue != zero && newValue == zero {
		return true
	}

	return false
}
