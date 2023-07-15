package cache

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v9"
	"github.com/khodemobin/golang_boilerplate/pkg/helper"
)

type client struct {
	rc  *redis.Client
	ctx context.Context
}

func New(rc *redis.Client) Cache {
	return &client{
		rc:  rc,
		ctx: context.Background(),
	}
}

func NewTest(t *testing.T) Cache {
	s := miniredis.RunT(t)
	r := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	return &client{
		rc:  r,
		ctx: context.Background(),
	}
}

func (r *client) Get(key string, defaultValue func() (string, error)) (string, error) {
	value, err := r.rc.Get(r.ctx, helper.ToMD5(key)).Result()
	if err == redis.Nil || value == "" {
		if defaultValue == nil {
			return "", nil
		}

		return defaultValue()
	}

	return value, err
}

func (r *client) Remember(key string, defaultValue func() (string, time.Duration, error)) (string, error) {
	value, err := r.rc.Get(r.ctx, helper.ToMD5(key)).Result()
	if err == redis.Nil {
		if defaultValue == nil {
			return "", nil
		}

		v, exp, err := defaultValue()
		if err != nil {
			return "", err
		}

		err = r.Set(key, v, exp)

		if err != nil {
			return "", err
		}

		return v, err
	}

	return value, err
}

func (r *client) Set(key string, value interface{}, expiration time.Duration) error {
	return r.rc.Set(r.ctx, helper.ToMD5(key), value, expiration).Err()
}

func (r *client) Delete(key string) error {
	return r.rc.Del(r.ctx, helper.ToMD5(key)).Err()
}

func (r *client) Pull(key string, defaultValue func() (string, error)) (string, error) {
	value, err := r.rc.Get(r.ctx, helper.ToMD5(key)).Result()
	if redis.Nil == err || value == "" {
		if defaultValue == nil {
			return "", nil
		}

		return defaultValue()
	}

	err = r.Delete(helper.ToMD5(key))

	if err != nil {
		return "", err
	}

	return value, err
}

func (r *client) Close() error {
	return r.rc.Close()
}
