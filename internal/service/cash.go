package service

import "petProject/internal/repository"

type CashService struct {
	CashRepository repository.CashRepository
}

func newCashService(cash *repository.CashRepository) *CashService {
	return &CashService{CashRepository: *cash}
}

func (s *CashService) CheckTokenExist(name, device string) error {
	return s.CashRepository.CheckTokenExist(name, device)
}

func (s *CashService) CheckUserAuthorized(token string) error {
	return s.CashRepository.CheckUserAuthorized(token)
}
