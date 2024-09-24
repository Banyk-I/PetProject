package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type CashRepository struct {
	redis *redis.Client
}

func NewCashRepository(redis *redis.Client) *CashRepository {
	return &CashRepository{
		redis: redis,
	}
}

func (r *CashRepository) CheckTokenExist(name, device string) error {
	ctx := context.TODO()

	key := name + ":" + device

	exists, err := r.redis.Exists(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("error checking if key exists: %v", err)
	}

	if exists > 0 {
		return fmt.Errorf("token already exists %v", key)
	}

	return nil
}

func (r *CashRepository) CheckUserAuthorized(token string) error {
	ctx := context.TODO()

	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = r.redis.Scan(ctx, cursor, "*", 10).Result()
		if err != nil {
			return fmt.Errorf("error scanning keys: %v", err)
		}

		for _, key := range keys {
			val, err := r.redis.Get(ctx, key).Result()
			if err != nil && err != redis.Nil {
				return fmt.Errorf("error getting key value: %v", err)
			}

			if val == token {
				return nil
			}
		}

		if cursor == 0 {
			// Ми пройшли всі ключі
			break
		}
	}

	return fmt.Errorf("token not authorized")
}
