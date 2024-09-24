package service

import "petProject/internal/repository"

type AdminService struct {
	adminRepository repository.AdminRepository
}

func newAdminService(repo *repository.AdminRepository) *AdminService {
	return &AdminService{
		adminRepository: *repo,
	}
}

func (s *AdminService) VerificationForAdmin(userID string) error {
	return s.adminRepository.VerificationForAdmin(userID)
}

func (s *AdminService) GetActiveAccessTokens() (map[string]string, error) {
	return s.adminRepository.GetActiveAccessTokens()
}

func (s *AdminService) LogoutUser(userLogout, device string) error {
	if device == "" {
		return s.adminRepository.LogoutUserAllDevices(userLogout)
	}
	return s.adminRepository.LogoutUserDevice(userLogout, device)
}
