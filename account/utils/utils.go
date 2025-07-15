package utils

type utils struct {
	CryptoUtils
	RandUtils
	AuthUtils
}

type Utils interface {
	CryptoUtils
	RandUtils
	AuthUtils
}

func NewUtilsConfig() Utils {
	return &utils{
		CryptoUtils: NewCryptoUtils(),
		RandUtils:   NewRandUtils(),
		AuthUtils:   NewAuthUtils(),
	}
}
