package testutils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"petProject/internal/repository"
	"testing"
)

func SetupAuth(t *testing.T) (*repository.AuthorizationRepository, func()) {
	// Налаштування підключення до MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatalf("failed to connect to MongoDB: %v", err)
	}

	// Вибір бази даних
	db := client.Database("users_test")

	// Налаштування підключення до Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Створення нового репозиторію
	repo := repository.NewAuthorizationRepositoryMongoDB(db, redisClient)

	// Функція для очищення ресурсів після тестування
	cleanup := func() {
		// Видалення бази даних MongoDB
		if err := client.Database("testdb").Drop(context.Background()); err != nil {
			t.Fatalf("failed to drop database: %v", err)
		}
		// Очищення бази даних Redis
		if err := redisClient.FlushDB(context.Background()).Err(); err != nil {
			t.Fatalf("failed to flush Redis DB: %v", err)
		}
	}

	return repo, cleanup
}

func SetupAdmin(t *testing.T) (*repository.AdminRepository, func()) {
	// Налаштування підключення до MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatalf("failed to connect to MongoDB: %v", err)
	}

	// Вибір бази даних
	db := client.Database("users_test")

	// Налаштування підключення до Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Створення нового репозиторію
	repo := repository.NewAdminRepository(db, redisClient)

	// Функція для очищення ресурсів після тестування
	cleanup := func() {
		// Видалення бази даних MongoDB
		if err := client.Database("testdb").Drop(context.Background()); err != nil {
			t.Fatalf("failed to drop database: %v", err)
		}
		// Очищення бази даних Redis
		if err := redisClient.FlushDB(context.Background()).Err(); err != nil {
			t.Fatalf("failed to flush Redis DB: %v", err)
		}
	}

	return repo, cleanup
}

func SetupCash(t *testing.T) (*repository.CashRepository, func()) {
	// Налаштування підключення до Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Створення нового репозиторію
	repo := repository.NewCashRepository(redisClient)

	// Функція для очищення ресурсів після тестування
	cleanup := func() {
		// Очищення бази даних Redis
		if err := redisClient.FlushDB(context.Background()).Err(); err != nil {
			t.Fatalf("failed to flush Redis DB: %v", err)
		}
	}

	return repo, cleanup
}
