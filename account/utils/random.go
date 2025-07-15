package utils

import (
	"math/rand"
	"time"
)

type RandUtils interface {
	GenerateSixDigitOtp() int
}

type randomUtils struct{}

func NewRandUtils() RandUtils {
	return &randomUtils{}
}
func (r *randomUtils) GenerateSixDigitOtp() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := 100000 + rng.Intn(900000)
	return otp
}
