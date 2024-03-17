package custommath

import (
	"crypto/rand"
	"math/big"

	log "github.com/sirupsen/logrus"
)

func GenerateRandomNumber(maxNumber int) int64 {
	randNumber, err := rand.Int(rand.Reader, big.NewInt(int64(maxNumber+1)))
	if err != nil {
		log.Error(err.Error())
		return int64(maxNumber)
	}
	return randNumber.Int64()
}

func GenerateRandomNumberByRange(minNumber int, maxNumber int) int {
	if minNumber >= maxNumber {
		log.Error("min should be less than max")
		return minNumber
	}

	return int(GenerateRandomNumber(maxNumber-minNumber)) + minNumber
}

