package ethutils

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"github.com/ethereum/go-ethereum/ethclient"
	redisV9 "github.com/redis/go-redis/v9"
)

func NewBlockCounter(name string, redisClient *redisV9.Client, ethClient *ethclient.Client) *BlockCounter {
	return &BlockCounter{name: name, redisClient: redisClient, ethClient: ethClient}
}

type BlockCounter struct {
	name        string
	redisClient *redisV9.Client
	ethClient   *ethclient.Client
}

func (t *BlockCounter) Current(ctx context.Context, minNumber int64) (int64, int64, error) {
	key := fmt.Sprintf("blockCounter:%s", t.name)

	// 当前最大
	blockNumber, err := t.ethClient.BlockNumber(ctx)
	if err != nil {
		return 0, 0, err
	}

	if minNumber > int64(blockNumber) {
		return 0, 0, fmt.Errorf("invalid min number")
		//minNumber = blockNumber
	}

	number, err := t.redisClient.Get(ctx, key).Int64()
	if errors.Is(err, redisV9.Nil) {
		t.redisClient.Set(ctx, key, minNumber, -1)
		return minNumber, int64(blockNumber), nil
	}

	if err != nil {
		slf.WithError(err).Errorw("Redis Get err", slf.String("key", key))
		return 0, 0, xerror.Wrap(err)
	}

	return number, int64(blockNumber), nil
}

func (t *BlockCounter) Incr(ctx context.Context) error {
	key := fmt.Sprintf("blockCounter:%s", t.name)

	_, err := t.redisClient.Incr(ctx, key).Result()
	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

func (t *BlockCounter) IncrBy(ctx context.Context, incrBy int64) error {
	key := fmt.Sprintf("blockCounter:%s", t.name)

	_, err := t.redisClient.IncrBy(ctx, key, incrBy).Result()
	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

func (t *BlockCounter) Decr(ctx context.Context) error {
	key := fmt.Sprintf("blockCounter:%s", t.name)

	_, err := t.redisClient.Decr(ctx, key).Result()
	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

func (t *BlockCounter) DecrBy(ctx context.Context, decrBy int64) error {
	key := fmt.Sprintf("blockCounter:%s", t.name)

	_, err := t.redisClient.DecrBy(ctx, key, decrBy).Result()
	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}
