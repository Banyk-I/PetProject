package testutils

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"petProject/internal/model"
	"testing"
)

func TestCreateUser_HappyPath(t *testing.T) {
	repo, clean := SetupAuth(t)
	defer clean()

	user := model.User{
		Name:     "maki",
		Password: "maki",
		Role:     "admin",
	}

	createdUser, err := repo.CreateUser(user)

	assert.NoError(t, err)
	assert.NotEqual(t, primitive.NilObjectID, createdUser.ID)
}

func TestGetUser(t *testing.T) {
	repo, clean := SetupAuth(t)
	defer clean()

	user := model.User{Name: "testuser", Password: "password123"}
	_, err := repo.CreateUser(user)
	assert.NoError(t, err)

	gotUser, err := repo.GetUser("testuser", "password123")
	assert.NoError(t, err)
	assert.Equal(t, "testuser", gotUser.Name)
}

func TestCheckIsUserExist(t *testing.T) {
	repo, clean := SetupAuth(t)
	defer clean()

	user := model.User{
		Name:     "maki",
		Password: "password123",
	}
	_, err := repo.CreateUser(user)
	assert.NoError(t, err)

	err = repo.CheckIsUserExist("maki")
	assert.Error(t, err)

	err = repo.CheckIsUserExist("nonexistentuser")
	assert.NoError(t, err)
}

func TestWriteTokenInRedis(t *testing.T) {
	repo, clean := SetupAuth(t)
	defer clean()

	err := repo.WriteTokenInRedis("testuser", "some-token", "device1")
	assert.NoError(t, err)

	token, err := repo.RedisClient.Get(context.Background(), "testuser:device1").Result()
	assert.NoError(t, err)
	assert.Equal(t, "some-token", token)
}

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, 10)
}
