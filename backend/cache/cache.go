package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheMissResolverFn func(key string) (interface{}, error)
type Cache struct {
	rbd           *redis.Client
	defaultExpiry time.Duration
}

var cache Cache

func Connect(address string, password string, defaultExpiry time.Duration) {
	rbd := redis.NewClient(&redis.Options{Addr: address, Password: password, DB: 0})
	cache = Cache{rbd: rbd, defaultExpiry: defaultExpiry}
}

func Set(ctx context.Context, key string, value interface{}, expiry ...time.Duration) error {
	var exp time.Duration
	if len(expiry) > 0 {
		exp = expiry[0]
	} else {
		exp = cache.defaultExpiry
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(value); err != nil {
		return err
	}

	return cache.rbd.Set(ctx, key, buf.Bytes(), exp).Err()
}

func Get(ctx context.Context, key string, value interface{}, callback CacheMissResolverFn) error {
	data, err := cache.rbd.Get(ctx, key).Bytes()

	if err == redis.Nil && callback != nil {
		// If Cache miss run the callback function
		newValue, err := callback(key)
		if err != nil {
			return err
		}

		err = Set(ctx, key, newValue)
		if err != nil {
			return err
		}

		// Conver and copy newValue to Buffer
		var newValueBuffer bytes.Buffer
		if err := gob.NewEncoder(&newValueBuffer).Encode(newValue); err != nil {
			return err
		}

		// Convert and copy newVaueBuffer to value
		if err := gob.NewDecoder(&newValueBuffer).Decode(value); err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	} else {
		buf := bytes.NewBuffer(data)

		dec := gob.NewDecoder(buf)

		if err := dec.Decode(value); err != nil {
			return err
		}
	}

	return nil
}

func Invalidate(ctx context.Context, key string) error {
	return cache.rbd.Del(ctx, key).Err()
}
