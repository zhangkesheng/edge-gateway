package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

func JwtEncode(identity *api.Identity, secret string) (string, error) {
	onError := func(err error) (string, error) {
		return "", errors.Wrap(err, "Jwt.Encode")
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nick":   identity.GetNick(),
		"avatar": identity.GetAvatar(),
		"email":  identity.GetEmail(),
		"source": identity.GetSource(),
	})
	tokenString, err := claims.SignedString([]byte(secret))
	if err != nil {
		return onError(err)
	}

	return tokenString, nil
}
