package authenticator

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	"otus-microservice-architecture/app/config"
)

const (
	JWTAuthSubject = "auth"
)

var (
	AuthorizationHeaderNotSetError = errors.New("authorization header not set")
	TokenBearerFormatError         = errors.New("authorization header format must be Bearer {token}")
	TokenBearerEmptyError          = errors.New("authorization token is empty")
	NotAuthorizedError             = errors.New("you're unauthorized")
	UnexpectedSigningMethod        = errors.New("unexpected signing method")
	InvalidTokenError              = errors.New("you're Unauthorized due to invalid token")
)

type Authenticator interface {
	Authenticate(request *http.Request) (int, error)
	GenerateJWTToken(user string) (string, error)
}

type Config struct {
	Log    *zap.Logger
	Config *config.AuthConfig
}

type Service struct {
	log    *zap.Logger
	config *config.AuthConfig
}

func New(config Config) *Service {
	return &Service{
		log:    config.Log,
		config: config.Config,
	}
}
