package account

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestSm(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()
	redisCli := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	sm := newRedisSessionManager(redisCli, 60*1000, "Test", "Test")
	ctx := context.Background()
	testUser := "TestUser123"
	testUser2 := "TestUser456"

	t.Run("New", func(t *testing.T) {
		token1, err := sm.New(ctx, testUser)
		assert.NoError(t, err)
		token2, err := sm.New(ctx, testUser2)
		assert.NoError(t, err)

		assert.NotEqual(t, token1.AccessToken, token2.AccessToken)
	})

	t.Run("Verify", func(t *testing.T) {
		token, err := sm.New(ctx, testUser)
		if !assert.NoError(t, err) {
			return
		}

		sub, err := sm.Verify(ctx, token.AccessToken)
		if assert.NoError(t, err) {
			assert.Equal(t, testUser, sub)
		}
	})

	t.Run("Verify error", func(t *testing.T) {
		sub, err := sm.Verify(ctx, "xxxxxxxxxxxxxxxx")
		if assert.Error(t, err) {
			assert.Empty(t, sub)
		}
	})

	t.Run("Clear", func(t *testing.T) {
		token, err := sm.New(ctx, testUser)
		if !assert.NoError(t, err) {
			return
		}

		err = sm.Clear(ctx, token.AccessToken)
		assert.NoError(t, err)
	})

	t.Run("Clear not exist token", func(t *testing.T) {
		err = sm.Clear(ctx, "ttt")
		assert.NoError(t, err)
	})

	t.Run("Refresh", func(t *testing.T) {
		token, err := sm.New(ctx, testUser)
		if !assert.NoError(t, err) {
			return
		}

		time.Sleep(time.Second)
		token2, err := sm.Refresh(ctx, token.AccessToken)
		if assert.NoError(t, err) {
			assert.NotEqual(t, token2.AccessToken, token.AccessToken)
		}
	})
}
