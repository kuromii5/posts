package service

type UserReg struct {
	username string `validate:"required,min=2"`
	email    string `validate:"required,email"`
	password string `validate:"required,min=8"`
}

func (s *Service) validateUserReg(username, email, password string) error {
	v := &UserReg{
		username: username,
		email:    email,
		password: password,
	}
	if err := s.v.Struct(v); err != nil {
		return err
	}
	return nil
}
