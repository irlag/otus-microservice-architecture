package password

import (
	"go.uber.org/zap"
)

const (
	SaltLength            = 6
	PBKDF2Iterations      = 1000
	PBKDF2KeyLength       = 16
	PasswordSaltSeparator = ":"
)

type PasswordHash interface {
	GetRandomSalt() (string, error)
	GeneratePasswordHashWithSalt(password string, salt string) (passwordHash string, err error)
	GeneratePasswordHashWithOutSalt(password string, salt string) (passwordHash string, err error)
	SplitPassword(passwordHash string) (string, string)
	CheckPassword(password string, passwordHash string) (bool, error)
}

type Config struct {
	Log *zap.Logger
}

type Service struct {
	log *zap.Logger
}

func New(config Config) *Service {
	return &Service{
		log: config.Log,
	}
}
