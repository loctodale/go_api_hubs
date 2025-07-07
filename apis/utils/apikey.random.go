package utils

import (
	"math/rand"
	"time"
)

type ApiKeyRandom interface {
	GenerateRandomApiKey(length int) string
}
type apiKeyRandom struct{}

func NewApiKeyRandom() ApiKeyRandom {
	return &apiKeyRandom{}
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (a *apiKeyRandom) GenerateRandomApiKey(length int) string {
	return StringWithCharset(length, charset)
}
