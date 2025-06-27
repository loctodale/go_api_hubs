package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type CryptoUtils interface {
	GetHash(key string) string
	GenerateSalt(length int) (string, error)
	HashPasswordSalt(password string, salt string) string
	MatchPassword(storeHash string, password string, salt string) bool
}
type cryptoUtils struct{}

func NewCryptoUtils() CryptoUtils {
	return &cryptoUtils{}
}

func (u *cryptoUtils) GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
func (u *cryptoUtils) GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}
func (u *cryptoUtils) HashPasswordSalt(password string, salt string) string {
	saltedPassword := password + salt

	hashPassword := sha256.Sum256([]byte(saltedPassword))
	return hex.EncodeToString(hashPassword[:])
}

func (u *cryptoUtils) MatchPassword(storeHash string, password string, salt string) bool {
	hashPassword := u.HashPasswordSalt(password, salt)
	return storeHash == hashPassword
}
