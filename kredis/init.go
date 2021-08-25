package kredis

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

var Subscribe *subscribe

type subscribe struct {
	DbRedis  *redis.Client
	channels map[string]*redis.PubSub
}

// Init 全局 Subscribe 初始化
func Init(ctx context.Context, client *redis.Client, channels map[string]func(ctx context.Context, pubSub *redis.PubSub)) {
	Subscribe = &subscribe{
		DbRedis:  client,
		channels: make(map[string]*redis.PubSub),
	}
	for channel, fc := range channels {
		ps := Subscribe.Sub(ctx, channel)
		Subscribe.channels[channel] = ps
		go fc(ctx, ps)
	}
}

// Close 消息订阅关闭
func Close() error {
	if Subscribe == nil {
		return nil
	}
	for _, pubSub := range Subscribe.channels {
		_ = pubSub.Close()
	}
	_ = Subscribe.DbRedis.Close()
	return nil
}

// Sub 订阅
func (s *subscribe) Sub(ctx context.Context, channel string) *redis.PubSub {
	var err error
	pubSub := s.DbRedis.Subscribe(ctx, channel)
	_, err = pubSub.Receive(ctx)
	if err != nil {
		return nil
	}
	return pubSub
}

// Pub 发布
func (s *subscribe) Pub(ctx context.Context, channel string, message []byte) error {
	var err error
	err = s.DbRedis.Publish(ctx, channel, message).Err()
	if err != nil {
		fmt.Printf("subscribe error: %s", err.Error())
		return err
	}
	return nil
}
