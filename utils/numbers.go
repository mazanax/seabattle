package utils

import (
	"crypto/rand"
	"math/big"
)

func RandomInt(below int64) (int64, error) {
	bigInt, err := rand.Int(rand.Reader, big.NewInt(below))
	if nil != err {
		return 0, err
	}
	return bigInt.Int64(), nil
}
