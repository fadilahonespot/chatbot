package cached

import (
	"context"
	"time"
)

type CacheWrapper interface {
	Set(ctx context.Context, key, value string, duration time.Duration) (err error)
	Get(ctx context.Context, key string) (value string, err error)
	Delete(ctx context.Context, key string) (err error)
}
