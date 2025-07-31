package cache

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestRedisCache(t *testing.T) {
	server, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	key := "testkey"
	value := "testvalue"

	conn := redis.NewClient(&redis.Options{Addr: server.Addr()})

	ctx := context.Background()
	mem := NewRedisCache(conn)
	timeout := 10 * time.Second
	err = mem.Set(ctx, key, value, timeout)
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
	err = mem.Del(ctx, "unknow-key")
	if err != nil {
		t.Fatalf("del error:%+v", err)
	}
}
