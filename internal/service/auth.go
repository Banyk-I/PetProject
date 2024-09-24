package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"petProject/internal/model"
	"petProject/internal/repository"
	"time"
)

const (
	tokenTTL = 15 * time.Minute
)

type authService struct {
	authorizationRepository *repository.AuthorizationRepository
	jwt.StandardClaims
}

func newAuthorizationService(repo *repository.AuthorizationRepository) *authService {
	return &authService{authorizationRepository: repo}
}

func (s *authService) CreateUser(user model.User) (model.User, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.authorizationRepository.CreateUser(user)
}

func (s *authService) GenerateToken(username, password string) (string, error) {
	user, err := s.authorizationRepository.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.TokenClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
		UserID:    user.ID.Hex(),
	})
	return token.SignedString([]byte(viper.GetString("secretKey")))
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(viper.GetString("JWT.salt"))))
}

func (s *authService) CheckIsUserExist(userName string) error {
	return s.authorizationRepository.CheckIsUserExist(userName)
}

func (s *authService) WriteTokenInRedis(username, token, device string) error {
	if device != "phone" && device != "laptop" {
		return fmt.Errorf("Device can be <<phone>> or <<laptop>> ")
	}
	return s.authorizationRepository.WriteTokenInRedis(username, token, device)
}
