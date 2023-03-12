package config

import (
	"time"
)

type AuthConfig struct {
	TokenTTL     time.Duration
	JWTSecretKey string
}
