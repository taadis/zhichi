package cache

import (
	"context"
	"sync"
	"time"
)

// check impl
var _ Cache = (*Memory)(nil)

// Memory local memory cache.
type Memory struct {
	sync.RWMutex
	//
	data map[string]*data
}

type data struct {
	Data interface{}
	//
	Expired time.Time
}

func NewMemoryCache() *Memory {
	m := new(Memory)
	m.data = make(map[string]*data)
	return m
}

func (m *Memory) Get(ctx context.Context, key string) interface{} {
	m.RLock()
	defer m.RUnlock()

	ret, ok := m.data[key]
	if !ok {
		return nil
	}
	if ret.Expired.Before(time.Now()) {
		m.del(ctx, key)
		return nil
	}

	return ret.Data
}

// Set key and value with expire time.
func (m *Memory) Set(ctx context.Context, key string, v interface{}, d time.Duration) error {
	m.Lock()
	defer m.Unlock()

	data := &data{
		Data:    v,
		Expired: time.Now().Add(d),
	}

	m.data[key] = data
	return nil
}

// Has check key exists.
func (m *Memory) Has(ctx context.Context, key string) bool {
	m.RLock()
	defer m.RUnlock()

	ret, ok := m.data[key]
	if !ok {
		return false
	}
	if ret.Expired.Before(time.Now()) {
		m.del(ctx, key)
		return false
	}

	return true
}

// Del delete key and value.
func (m *Memory) Del(ctx context.Context, key string) error {
	return m.del(ctx, key)
}

func (m *Memory) del(ctx context.Context, key string) error {
	m.Lock()
	defer m.Unlock()

	delete(m.data, key)

	return nil
}
