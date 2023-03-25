package password

func (p *Service) CheckPassword(password string, passwordHash string) (bool, error) {
	passHash, salt := p.SplitPassword(passwordHash)

	passGenHash, err := p.GeneratePasswordHashWithOutSalt(password, salt)
	if err != nil {
		return false, err
	}

	return passHash == passGenHash, nil
}
