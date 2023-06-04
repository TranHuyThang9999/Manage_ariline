package helper

import (
	"btl/config"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

func GetValueCache(key string) ([]byte, error) {
	val, err := ClientRedis().Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return nil, nil // Không tìm thấy giá trị trong cache
	} else if err != nil {
		return nil, err // Xử lý lỗi khi truy vấn cache
	}
	return val, nil
}

func SetValueCache(ctx context.Context, key string, value interface{}, expiration time.Duration) ([]byte, error) {
	client := ClientRedis()

	// Kiểm tra xem khóa đã tồn tại trong Redis chưa
	exists, err := client.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if exists == 1 {
		// Khóa đã tồn tại - lấy dữ liệu từ Redis
		val, err := client.Get(ctx, key).Bytes()
		if err != nil {
			return nil, err
		}
		return val, nil
	}
	// Chuyển đổi giá trị thành bytes
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	// Lưu dữ liệu vào Redis cache
	err = client.Set(ctx, key, valueBytes, expiration).Err() // Lưu trong 5 phút
	if err != nil {
		return nil, err
	}

	return valueBytes, nil
}

func ClientRedis() *redis.Client {
	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		return nil
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       config.Redis.DB,
	})
	return rdb
}
