package testutils

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"petProject/internal/model"
	"testing"
)

func TestVerificationForAdmin(t *testing.T) {
	repo, clean := SetupAdmin(t)
	defer clean()

	userOne := model.User{
		ID:       primitive.NewObjectID(),
		Name:     "maki",
		Password: "maki",
		Role:     "admin",
	}

	userTwo := model.User{
		ID:       primitive.NewObjectID(),
		Name:     "maki",
		Password: "maki",
		Role:     "",
	}

	r, c := SetupAuth(t)
	defer c()

	_, _ = r.CreateUser(userOne)
	_, _ = r.CreateUser(userTwo)

	err := repo.VerificationForAdmin(userOne.ID.Hex())
	assert.NoError(t, err)

	err = repo.VerificationForAdmin(userTwo.ID.Hex())
	assert.Error(t, err)
}

func TestGetActiveAccessTokens(t *testing.T) {
	repoAdmin, clean := SetupAdmin(t)
	defer clean()
	repoAuth, cleanup := SetupAuth(t)
	defer cleanup()

	_ = repoAuth.WriteTokenInRedis("testuser", "some-token", "phone")
	_ = repoAuth.WriteTokenInRedis("testuser2", "some-token2", "laptop")

	tokenMap, err := repoAdmin.GetActiveAccessTokens()
	assert.NoError(t, err)

	// Очікуваний результат
	expectedMap := map[string]string{
		"testuser:phone":   "some-token",
		"testuser2:laptop": "some-token2",
	}

	// Перевіряємо, що отримана мапа відповідає очікуваній
	assert.Equal(t, expectedMap, tokenMap)
}

func TestLogoutUserDevice(t *testing.T) {
	repoAdmin, clean := SetupAdmin(t)
	defer clean()
	repoAuth, cleanup := SetupAuth(t)
	defer cleanup()

	_ = repoAuth.WriteTokenInRedis("testuser", "some-token", "phone")
	_ = repoAuth.WriteTokenInRedis("testuser2", "some-token2", "laptop")

	err := repoAdmin.LogoutUserDevice("testuser", "phone")
	assert.NoError(t, err)

	tokenMap, err := repoAdmin.GetActiveAccessTokens()
	assert.NoError(t, err)

	expectedMap := map[string]string{
		"testuser2:laptop": "some-token2",
	}

	// Перевіряємо, що отримана мапа відповідає очікуваній
	assert.Equal(t, expectedMap, tokenMap)
}
func TestLogoutUserAllDevices(t *testing.T) {
	repoAdmin, clean := SetupAdmin(t)
	defer clean()
	repoAuth, cleanup := SetupAuth(t)
	defer cleanup()

	err := repoAdmin.LogoutUserAllDevices("testuser")
	assert.Error(t, err)

	_ = repoAuth.WriteTokenInRedis("testuser", "some-token", "phone")
	_ = repoAuth.WriteTokenInRedis("testuser", "some-token2", "laptop")

	err = repoAdmin.LogoutUserAllDevices("testuser")
	assert.NoError(t, err)

	tokenMap, err := repoAdmin.GetActiveAccessTokens()
	assert.NoError(t, err)

	expectedMap := map[string]string{}
	assert.Equal(t, expectedMap, tokenMap)

}
