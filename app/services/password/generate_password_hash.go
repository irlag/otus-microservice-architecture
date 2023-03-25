package password

import (
	"crypto/sha1"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

func (p *Service) GeneratePasswordHashWithSalt(password string, salt string) (passwordHash string, err error) {
	if salt == "" {
		salt, err = p.GetRandomSalt()
		if err != nil {
			return "", err
		}
	}

	hashedPassword := pbkdf2.Key([]byte(password), []byte(salt), PBKDF2Iterations, PBKDF2KeyLength, sha1.New)
	passwordHash = base64.StdEncoding.EncodeToString(hashedPassword) +
		PasswordSaltSeparator +
		salt

	return passwordHash, nil
}

func (p *Service) GeneratePasswordHashWithOutSalt(password string, salt string) (passwordHash string, err error) {
	hashedPassword := pbkdf2.Key([]byte(password), []byte(salt), PBKDF2Iterations, PBKDF2KeyLength, sha1.New)
	passwordHash = base64.StdEncoding.EncodeToString(hashedPassword)

	return passwordHash, nil
}
