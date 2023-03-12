package password

import (
	"strings"
)

func (p *Service) SplitPassword(passwordHash string) (password string, salt string) {
	pass := strings.Split(passwordHash, PasswordSaltSeparator)

	return pass[0], pass[1]
}
