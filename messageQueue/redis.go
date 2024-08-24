package messageQueue

import (
	"context"
	"fmt"

	"github.com/gambitier/gocomm/config"
	"github.com/go-redis/redis/v8"
)

// `RedisMessageQueue` implements `MessageQueue` interface using Redis.
type RedisMessageQueue struct {
	client *redis.Client
}

func NewRedisMessageQueue(conf *config.Conf) (*RedisMessageQueue, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
	})

	// Ping the Redis server to ensure connection is successful
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &RedisMessageQueue{client: rdb}, nil
}

// Publish publishes a message to a channel in Redis.
func (q *RedisMessageQueue) Publish(channel string, message []byte) error {
	ctx := context.Background()
	err := q.client.Publish(ctx, channel, message).Err()
	return err
}

func (q *RedisMessageQueue) Subscribe(channel string, handler MessageHandler) error {
	ctx := context.Background()
	pubsub := q.client.Subscribe(ctx, channel)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		return err
	}

	go func() {
		defer pubsub.Close()
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				break
			}
			handler([]byte(msg.Payload))
		}
	}()

	return nil
}

func (q *RedisMessageQueue) Close() error {
	return q.client.Close()
}
