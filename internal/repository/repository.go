package repository

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	mongoDB       *mongo.Database
	redisClient   *redis.Client
	Authorization *AuthorizationRepository
	Admin         *AdminRepository
	Cash          *CashRepository
}

const usersDB = "users"

func NewRepository(mongoDB *mongo.Database, redisClient *redis.Client) *Repository {
	return &Repository{
		mongoDB:       mongoDB,
		redisClient:   redisClient,
		Authorization: NewAuthorizationRepositoryMongoDB(mongoDB, redisClient),
		Admin:         NewAdminRepository(mongoDB, redisClient),
		Cash:          NewCashRepository(redisClient),
	}
}
