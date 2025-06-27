package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AuthUtils interface {
	GenerateKey() (privateKey *rsa.PrivateKey, publicToken *rsa.PublicKey, err error)
	GenerateToken(username string, privateKey *rsa.PrivateKey, userId string) (accessToken string, refreshToken string, err error)
	PublicKeyToString(publicKey *rsa.PublicKey) (pubPEMString string, err error)
	CreateConsumer(mail string, userId string) error
	AddConsumerToACLGroup(account string, groupName string) error
	AddConsumerJWTCredential(account string, publicKey string, userId string) error
}
type authUtils struct {
}

func NewAuthUtils() AuthUtils {
	return &authUtils{}
}

func (authUtils *authUtils) GenerateKey() (privateKey *rsa.PrivateKey, publicToken *rsa.PublicKey, err error) {
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func (authUtils authUtils) GenerateToken(username string, privateKey *rsa.PrivateKey, userId string) (accessToken string, refreshToken string, err error) {
	timeExAccess, err := time.ParseDuration("1h")
	timeExRefresh, err := time.ParseDuration("72h")

	if err != nil {
		return "", "", err
	}

	accessClaims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(timeExAccess).Unix(),
	}
	refreshClaims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(timeExRefresh).Unix(),
	}

	access := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims)
	refresh := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)

	accessToken, err = access.SignedString(privateKey)
	refreshToken, err = refresh.SignedString(privateKey)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (authUtils authUtils) PublicKeyToString(publicKey *rsa.PublicKey) (pubPEMString string, err error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}

	// Create a PEM block
	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	// Convert []byte to string
	pubPEMString = string(pubPEM)
	return pubPEMString, nil
}

func (a *authUtils) CreateConsumer(mail string, userId string) error {
	form := url.Values{}
	form.Set("username", mail)
	form.Set("custom_id", userId)

	resp, err := http.Post("http://kong:8001/consumers/", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Status Code Register Account:", resp.StatusCode)
	fmt.Println("Response Body Register Account:", string(body))
	return nil
}

func (a *authUtils) AddConsumerToACLGroup(account string, groupName string) error {
	aclGroup := url.Values{}
	aclGroup.Add("group", groupName)
	resp, err := http.Post(fmt.Sprintf("http://kong:8001/consumers/%s/acls", account), "application/x-www-form-urlencoded", strings.NewReader(aclGroup.Encode()))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Status Code Register ACL:", resp.StatusCode)
	fmt.Println("Response Body Register ACL:", resp.Body)
	return nil
}

func (a *authUtils) AddConsumerJWTCredential(account string, publicKey string, userId string) error {
	url := fmt.Sprintf("http://kong:8001/consumers/%s/jwt", account)

	jwtData := map[string]string{
		"algorithm":      "RS256",
		"rsa_public_key": publicKey,
		"key":            account,
	}
	jsonData, err := json.Marshal(jwtData)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Status Code AddConsumerJWTCredential:", resp.StatusCode)
	fmt.Println("Response Body AddConsumerJWTCredential:", resp.Body)
	fmt.Println("JWT credential created successfully")

	return nil
}
