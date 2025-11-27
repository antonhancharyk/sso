package security

import (
	"crypto/rsa"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type RSA struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func MustLoad() (RSA, error) {
	privateKey, err := LoadPrivateKey("./config/rsa/private_key.pem")
	if err != nil {
		return RSA{}, err
	}

	publicKey, err := LoadPublicKey("./config/rsa/public_key.pem")
	if err != nil {
		return RSA{}, err
	}

	return RSA{PrivateKey: privateKey, PublicKey: publicKey}, nil
}

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateKeyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	publicKeyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
