package custommath

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func GenerateRandomNumber(maxNumber int) (int64, error) {
	randNumber, err := rand.Int(rand.Reader, big.NewInt(int64(maxNumber+1)))
	if err != nil {
		return int64(maxNumber), err
	}
	return randNumber.Int64(), nil
}

func GenerateRandomNumberByRange(minNumber int, maxNumber int) (int, error) {
	if minNumber >= maxNumber {
		return minNumber, errors.New("min should be less than max")
	}

	randomNumber, err := GenerateRandomNumber(maxNumber - minNumber)

	return int(randomNumber) + minNumber, err
}
