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
type utils struct{}

func NewUtils() CryptoUtils {
	return &utils{}
}

func (u *utils) GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
func (u *utils) GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}
func (u *utils) HashPasswordSalt(password string, salt string) string {
	saltedPassword := password + salt

	hashPassword := sha256.Sum256([]byte(saltedPassword))
	return hex.EncodeToString(hashPassword[:])
}

func (u *utils) MatchPassword(storeHash string, password string, salt string) bool {
	hashPassword := u.HashPasswordSalt(password, salt)
	return storeHash == hashPassword
}
