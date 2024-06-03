package config

import (
	"github.com/antongoncharik/sso/internal/entity"
	"github.com/antongoncharik/sso/pkg/utilities"
)

func MustLoad() (entity.RSAKeys, error) {
	privateKey, err := utilities.LoadPrivateKey("../../config/rsa/private_key.pem")
	if err != nil {
		return entity.RSAKeys{}, err
	}

	publicKey, err := utilities.LoadPublicKey("../../config/rsa/public_key.pem")
	if err != nil {
		return entity.RSAKeys{}, err
	}

	return entity.RSAKeys{PrivateKey: privateKey, PublicKey: publicKey}, nil
}
