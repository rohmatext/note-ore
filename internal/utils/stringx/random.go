package stringx

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Random(length int) (string, error) {
	result := make([]rune, length)
	runes := []rune(charset)
	x := int64(len(runes))
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(x))
		if err != nil {
			return "", errors.New("error creating random number")
		}
		result[i] = runes[num.Int64()]
	}
	return string(result), nil
}
