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

func RandomInt8(below int8) (int8, error) {
	result, err := RandomInt(int64(below))

	if err != nil {
		return 0, err
	}

	return int8(result), nil
}

func ContainsBit(mask uint64, bit int) bool {
	return mask&uint64(bit) == uint64(bit)
}
