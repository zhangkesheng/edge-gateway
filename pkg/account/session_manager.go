package account

import (
	"context"
	"reflect"
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
	Refresh(ctx context.Context, token string) error
	Verify(ctx context.Context, token string) (string, error)
	Clear(ctx context.Context, token string) error
}

type redisSM struct {
	cli       *redis.Client
	expiresIn int64
	secret    string
	issuer    string
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

	if err := r.cli.Set(tokenString, sub, time.Duration(r.expiresIn)*time.Second).Err(); err != nil {
		return onError(err)
	}

	return token, nil
}

func (r *redisSM) Refresh(ctx context.Context, token string) error {
	onError := func(err error) error {
		return errors.Wrap(err, "RedisSM.Refresh")
	}
	if err := r.cli.Expire(token, time.Duration(r.expiresIn)*time.Second).Err(); err != nil {
		return onError(err)
	}
	return nil
}

func (r *redisSM) Verify(ctx context.Context, token string) (string, error) {
	onError := func(err error) (string, error) {
		return "", errors.Wrap(err, "RedisSM.Verify")
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

	claims := jwtToken.Claims
	switch reflect.TypeOf(claims).Name() {
	case "MapClaims":
		return claims.(jwt.MapClaims)["sub"].(string), nil
	default:
		return claims.(jwt.StandardClaims).Subject, nil
	}
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
