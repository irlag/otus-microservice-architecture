package password

import (
	"crypto/rand"
	"encoding/hex"
)

func (p *Service) GetRandomSalt() (string, error) {
	var salt = make([]byte, SaltLength)

	_, err := rand.Read(salt[:])
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}
