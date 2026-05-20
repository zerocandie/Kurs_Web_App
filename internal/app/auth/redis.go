package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// RedisClient обертка над redis.Client
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient создает новое подключение к Redis
func NewRedisClient(host string, port string, password string, db int) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	// Проверка подключения
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &RedisClient{client: client}, nil
}

// Set устанавливает значение по ключу с TTL
func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get получает значение по ключу
func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Delete удаляет ключ
func (r *RedisClient) Delete(key string) error {
	return r.client.Del(ctx, key).Err()
}

// Exists проверяет существование ключа
func (r *RedisClient) Exists(key string) bool {
	result, _ := r.client.Exists(ctx, key).Result()
	return result > 0
}

// Close закрывает подключение
func (r *RedisClient) Close() error {
	return r.client.Close()
}
