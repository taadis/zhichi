package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

// check impl
var _ Cache = (*Memcache)(nil)

// Memcache struct contains *memcache.Client
type Memcache struct {
	conn *memcache.Client
}

func NewMemcache(server ...string) *Memcache {
	return &Memcache{
		memcache.New(server...),
	}
}

// Get return cached value.
func (c *Memcache) Get(ctx context.Context, key string) interface{} {
	item, err := c.conn.Get(key)
	if err != nil {
		return nil
	}

	var result interface{}
	err = json.Unmarshal(item.Value, &result)
	if err != nil {
		return nil
	}

	return result
}

// Set key and value with expire time.
func (c *Memcache) Set(ctx context.Context, key string, v interface{}, d time.Duration) error {
	buf, err := json.Marshal(v)
	if err != nil {
		return err
	}

	item := &memcache.Item{
		Key:        key,
		Value:      buf,
		Expiration: int32(d / time.Second),
	}

	return c.conn.Set(item)
}

// Has check key exists.
func (c *Memcache) Has(ctx context.Context, key string) bool {
	_, err := c.conn.Get(key)
	return err == nil
}

// Del delete key and value.
func (c *Memcache) Del(ctx context.Context, key string) error {
	return c.conn.Delete(key)
}
