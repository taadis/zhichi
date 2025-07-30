package cache

import (
	"context"
	"time"
)

// Cache interface
type Cache interface {
	Get(ctx context.Context, key string) interface{}
	Set(ctx context.Context, key string, value interface{}, timeout time.Duration) error
	Has(ctx context.Context, key string) bool
	Del(ctx context.Context, key string) error
}
