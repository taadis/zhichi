package cache

import (
	"context"
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

func TestMemcache(t *testing.T) {
	key := "testkey"
	value := "testvalue"

	ctx := context.Background()
	mem := NewMemcache("127.0.0.1:11211")
	timeout := 10 * time.Second
	err := mem.Set(ctx, key, value, timeout)
	if err != nil {
		t.Fatal(err)
	}

	// has
	if !mem.Has(ctx, key) {
		t.Fatal("not found key")
	}
	has := mem.Has(ctx, "not-found-key")
	if has != false {
		t.Fatalf("want false but got %v", has)
	}

	// get
	val := mem.Get(ctx, key).(string)
	if val != value {
		t.Fatalf("want %s but got %s", value, val)
	}
	nilval := mem.Get(ctx, "not-found-key")
	if nilval != nil {
		t.Fatalf("want nil but got %v", nilval)
	}

	// del
	err = mem.Del(ctx, key)
	if err != nil {
		t.Fatalf("del error:%+v", err)
	}
	err = mem.Del(ctx, "del-key")
	if err != memcache.ErrCacheMiss {
		t.Fatalf("del error:%+v", err)
	}
}
