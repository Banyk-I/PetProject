package service

import (
	"petProject/internal/repository"
)

type Service struct {
	repository           *repository.Repository
	AuthorizationService *authService
	AdminService         *AdminService
	CashService          *CashService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		repository:           repository,
		AuthorizationService: newAuthorizationService(repository.Authorization),
		AdminService:         newAdminService(repository.Admin),
		CashService:          newCashService(repository.Cash),
	}
}
