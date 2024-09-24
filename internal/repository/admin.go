package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"petProject/internal/model"
)

type AdminRepository struct {
	mongoClient *mongo.Collection
	redisClient *redis.Client
}

func NewAdminRepository(mongoDB *mongo.Database, redisClient *redis.Client) *AdminRepository {
	return &AdminRepository{
		mongoClient: mongoDB.Collection(usersDB),
		redisClient: redisClient,
	}
}

func (r *AdminRepository) VerificationForAdmin(userID string) error {
	ctx := context.TODO()

	// Підготовка фільтра для пошуку за ID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID format")
	}
	filter := bson.M{"_id": objID}

	// Виконання запиту до MongoDB
	var user model.User
	err = r.mongoClient.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user with ID %s not found", userID)
		}
		return fmt.Errorf("error finding user: %v", err)
	}

	// Перевірка ролі користувача
	if user.Role != "admin" {
		return fmt.Errorf("user with ID %s is not an admin", userID)
	}

	return nil
}

func (r *AdminRepository) GetActiveAccessTokens() (map[string]string, error) {
	ctx := context.TODO()

	keys := r.redisClient.Keys(ctx, "*").Val()
	tokenMap := make(map[string]string)

	for _, key := range keys {
		val, err := r.redisClient.Get(ctx, key).Result()
		if err != nil {
			return nil, fmt.Errorf("error getting value for key %s: %v", key, err)
		}
		tokenMap[key] = val
	}

	return tokenMap, nil
}

func (r *AdminRepository) LogoutUserDevice(userLogout, device string) error {
	ctx := context.TODO()

	key := userLogout + ":" + device

	exists, err := r.redisClient.Exists(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("error checking if key exists: %v", err)
	}

	if exists > 0 {
		_, err := r.redisClient.Del(ctx, key).Result()
		if err != nil {
			return fmt.Errorf("error deleting key: %v", err)
		}
	} else {
		return fmt.Errorf("user %s not found in active sessions", userLogout)
	}

	return nil
}

func (r *AdminRepository) LogoutUserAllDevices(userLogout string) error {
	ctx := context.TODO()

	// Зчитуємо всі ключі з Redis, які починаються з `userLogout:`
	pattern := userLogout + ":*"
	keys, err := r.redisClient.Keys(ctx, pattern).Result()
	if err != nil {
		return fmt.Errorf("failed to fetch keys from Redis: %w", err)
	}

	// Якщо немає жодного ключа, повертаємо помилку
	if len(keys) == 0 {
		return fmt.Errorf("no active sessions found for user %s", userLogout)
	}

	// Видаляємо кожен ключ з Redis
	for _, key := range keys {
		err := r.redisClient.Del(ctx, key).Err()
		if err != nil {
			return fmt.Errorf("failed to delete key %s from Redis: %w", key, err)
		}
	}

	return nil
}
