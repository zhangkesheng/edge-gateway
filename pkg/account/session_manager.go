package account

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type Token struct {
	AccessToken string
	RefreshToke string
	ExpiresIn   int64
}

type SessionManager interface {
	New(ctx context.Context, sub string) (*Token, error)
	NewWithIdentity(ctx context.Context, identity map[string]interface{}) (*Token, error)
	Refresh(ctx context.Context, token string) (*Token, error)
	Verify(ctx context.Context, token string) (map[string]interface{}, error)
	Clear(ctx context.Context, token string) error
}

type redisSM struct {
	cli       *redis.Client
	expiresIn int64
	secret    string
	issuer    string
}

func (r *redisSM) NewWithIdentity(ctx context.Context, identity map[string]interface{}) (*Token, error) {
	onError := func(err error) (*Token, error) {
		return nil, errors.Wrap(err, "RedisSM.NewWithIdentity")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(identity))

	tokenString, err := claims.SignedString([]byte(r.secret))
	if err != nil {
		return onError(err)
	}

	token := &Token{
		AccessToken: tokenString,
		ExpiresIn:   r.expiresIn,
	}

	if err := r.cli.Set(tokenString, identity, time.Duration(r.expiresIn)*time.Second).Err(); err != nil {
		return onError(err)
	}

	return token, nil
}

func (r *redisSM) New(ctx context.Context, sub string) (*Token, error) {
	onError := func(err error) (*Token, error) {
		return nil, errors.Wrap(err, "RedisSM.New")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    r.issuer,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Unix() + r.expiresIn,
		Subject:   sub,
	})

	tokenString, err := claims.SignedString([]byte(r.secret))
	if err != nil {
		return onError(err)
	}

	token := &Token{
		AccessToken: tokenString,
		ExpiresIn:   r.expiresIn,
	}

	if err := r.cli.Set(tokenString, map[string]interface{}{"sub": sub}, time.Duration(r.expiresIn)*time.Second).Err(); err != nil {
		return onError(err)
	}

	return token, nil
}

func (r *redisSM) Refresh(ctx context.Context, token string) (*Token, error) {
	onError := func(err error) (*Token, error) {
		return nil, errors.Wrap(err, "RedisSM.Refresh")
	}

	sub, err := r.Verify(ctx, token)
	if err != nil {
		return onError(err)
	}

	if err := r.Clear(ctx, token); err != nil {
		return onError(err)
	}

	newToken, err := r.NewWithIdentity(ctx, sub)
	if err != nil {
		return onError(err)
	}

	return newToken, nil
}

func (r *redisSM) Verify(ctx context.Context, token string) (map[string]interface{}, error) {
	onError := func(err error) (map[string]interface{}, error) {
		return nil, errors.Wrap(err, "RedisSM.Verify")
	}

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.secret), nil
	})
	if err != nil {
		return onError(err)
	}

	if err := r.cli.Get(token).Err(); err != nil {
		return onError(err)
	}

	return jwtToken.Claims.(jwt.MapClaims), nil
}

func (r *redisSM) Clear(ctx context.Context, token string) error {
	onError := func(err error) error {
		return errors.Wrap(err, "RedisSM.Clear")
	}
	if err := r.cli.Del(token).Err(); err != nil {
		return onError(err)
	}
	return nil
}

func newRedisSessionManager(redisCli *redis.Client, expiresIn int64, secret, issuer string) SessionManager {
	return &redisSM{
		cli:       redisCli,
		expiresIn: expiresIn,
		secret:    secret,
		issuer:    issuer,
	}
}
