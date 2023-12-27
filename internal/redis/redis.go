package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/santoadji21/santoadji21-go-fiber-chat/pkg/config"
)

var RDB *redis.Client

func ConnectRedis(cfg config.Config) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     cfg.RDHost + ":" + cfg.RDPort,
		Password: cfg.RDPassword,
		DB:       0,
	})

	// Testing the connection
	ctx := context.Background()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func GetRedis() *redis.Client {
	return RDB
}
