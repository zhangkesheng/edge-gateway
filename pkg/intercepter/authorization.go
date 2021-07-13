package intercepter

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zhangkesheng/edge-gateway/api/v1"
	"github.com/zhangkesheng/edge-gateway/pkg/types"
)

func Authorize(account api.AccountServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetToken(c)
		if token == "" {
			types.CtxError(c, http.StatusUnauthorized, errors.New("Not authorize "))
			return
		}

		ctx := c.Request.Context()
		if resp, err := account.Verify(ctx, &api.VerifyRequest{
			Token: token,
		}); err != nil {
			types.CtxError(c, http.StatusForbidden, err)
			return
		} else {
			c.Set("x-user-sub", resp.GetSub())
			c.Request.Header.Set("x-user-sub", resp.GetSub())
			c.Next()
		}
	}
}

func GetToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")

	if strings.HasPrefix(token, "Bearer ") {
		token = strings.Replace(token, "Bearer ", "", -1)
	}

	return strings.TrimSpace(token)
}
