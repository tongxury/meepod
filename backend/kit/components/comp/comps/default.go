package comps

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	redisV9 "github.com/redis/go-redis/v9"
	"time"
)

type DefaultProofStore struct{}

func (p *DefaultProofStore) ShouldSend(ctx context.Context, dest string) (bool, error) {
	set, err := comp.SDK().Redis().SetNX(ctx, fmt.Sprintf("auth_code.limit.%s", dest),
		"1", 1*time.Minute).Result()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return set, nil
}

func (p *DefaultProofStore) Save(ctx context.Context, dest, proof string) error {
	if _, err := comp.SDK().Redis().Set(ctx, fmt.Sprintf("auth_code.%s", dest),
		proof, 5*time.Minute).Result(); err != nil {

		return xerror.Wrap(err)
	}

	return nil
}
func (p *DefaultProofStore) Find(ctx context.Context, dest string) (string, error) {
	storeCode, err := comp.SDK().Redis().Get(ctx, fmt.Sprintf("auth_code.%s", dest)).Result()
	if err != nil && !errors.Is(err, redisV9.Nil) {
		return "", xerror.Wrap(err)
	}

	return storeCode, nil
}
