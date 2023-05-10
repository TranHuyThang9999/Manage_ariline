package cache

import (
	"btl/config"
	"btl/infrastructure/model"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheWrapper interface {
	GetAllTicket(ctx context.Context) ([]*model.Booking, error)
}

type RedisCache struct {
	cache       CacheWrapper
	redisClient *redis.Client
	ttl         time.Duration
}

func NewRedisCache(cache CacheWrapper, redisClient *redis.Client, ttl time.Duration) *RedisCache {

	return &RedisCache{
		cache:       cache,
		redisClient: redisClient,
		ttl:         ttl,
	}
}
func NewRedisClient(configFile string) *redis.Client {
	config, err := config.LoadConfig(configFile)
	if err != nil {
		return nil
	}

	return redis.NewClient(&redis.Options{
		Addr: config.Redis.Addr,
		DB:   config.Redis.DB,
	})
}

func (c *RedisCache) GetAllTicket(ctx context.Context) ([]*model.Booking, error) {
	cacheKey := "all_tickets"
	var tickets []*model.Booking
	data, err := c.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		err = json.Unmarshal([]byte(data), &tickets)
		if err != nil {
			return nil, err
		}
		return tickets, nil
	}
	if err != redis.Nil {
		return nil, err
	}
	tickets, err = c.cache.GetAllTicket(ctx)
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(tickets)
	if err != nil {
		return nil, err
	}
	err = c.redisClient.Set(ctx, cacheKey, jsonData, c.ttl).Err()
	if err != nil {
		return nil, err
	}
	return tickets, nil
}
