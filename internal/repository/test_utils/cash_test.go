package testutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckTokenExist(t *testing.T) {
	repoCash, clean := SetupCash(t)
	defer clean()

	repoAuth, cleanAuth := SetupAuth(t)
	defer cleanAuth()

	err := repoCash.CheckTokenExist("testuser", "phone")
	assert.NoError(t, err)

	_ = repoAuth.WriteTokenInRedis("testuser", "some-thing", "phone")
	err = repoCash.CheckTokenExist("testuser", "phone")
	assert.Error(t, err)

}

func TestCheckUserAuthorized(t *testing.T) {
	repoCash, clean := SetupCash(t)
	defer clean()

	repoAuth, cleanAuth := SetupAuth(t)
	defer cleanAuth()

	// Додавання токену в Redis
	token := "authorized-token"
	err := repoAuth.WriteTokenInRedis("testuser", token, "phone")
	assert.NoError(t, err, "expected no error when writing token")

	// Перевірка, що токен авторизований
	err = repoCash.CheckUserAuthorized(token)
	assert.NoError(t, err, "expected token to be authorized")

	// Перевірка, що неавторизований токен викликає помилку
	err = repoCash.CheckUserAuthorized("unauthorized-token")
	assert.Error(t, err, "expected error for unauthorized token")
	assert.Contains(t, err.Error(), "token not authorized", "unexpected error message")
}
