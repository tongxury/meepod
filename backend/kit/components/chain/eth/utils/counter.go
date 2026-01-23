package ethutils

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	redisV9 "github.com/redis/go-redis/v9"
)

func NewCounter(name string, initialValue int64, redisClient *redisV9.Client) *Counter {
	return &Counter{name: name, initialValue: initialValue, redisClient: redisClient}
}

type Counter struct {
	name         string
	initialValue int64
	redisClient  *redisV9.Client
}

func (t *Counter) Current(ctx context.Context) (int64, error) {
	key := fmt.Sprintf("counter_:%s", t.name)

	number, err := t.redisClient.Get(ctx, key).Int64()

	if err != nil {
		if errors.Is(err, redisV9.Nil) {
			_, err := t.redisClient.Set(ctx, key, t.initialValue, -1).Result()
			if err != nil {
				return 0, xerror.Wrap(err)
			}
			number = t.initialValue
		} else {
			return 0, xerror.Wrap(err)
		}
	}

	return number, nil
}

func (t *Counter) Incr(ctx context.Context) error {
	key := fmt.Sprintf("counter_:%s", t.name)

	_, err := t.redisClient.Incr(ctx, key).Result()
	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

func (t *Counter) Decr(ctx context.Context) error {
	key := fmt.Sprintf("counter_:%s", t.name)

	_, err := t.redisClient.Decr(ctx, key).Result()
	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

func (t *Counter) DecrBy(ctx context.Context, decrBy int64) error {
	key := fmt.Sprintf("counter_:%s", t.name)

	_, err := t.redisClient.DecrBy(ctx, key, decrBy).Result()
	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}
