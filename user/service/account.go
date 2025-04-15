package service

type AccountService interface {
	Register() error
}

type accountService service

func (s *accountService) Register() error {
	return nil
}
