package oauth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

const (
	testSessionKey = "testSessionKey"
)

func TestNewMiniProgram(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("js_code")
		if code == "test" {
			_, _ = w.Write([]byte(fmt.Sprintf(`{"errcode":0,"errmsg":"ok","openid":"%s","session_key":"%s","unionid":"%s"}`, testOpenId, testSessionKey, testUnionId)))
		} else {
			_, _ = w.Write([]byte(`{"errcode":111,"errmsg":"error"}`))
		}
	}))

	defer ts.Close()

	config := config{
		clientId:       "test",
		secret:         "test",
		accessTokenUrl: ts.URL + "/sns/jscode2session",
	}
	cli := NewOauthCli(MiniProgram, config)

	ctx := context.Background()

	t.Run("Token", func(t *testing.T) {
		resp, err := cli.AccessToken(ctx, &api.AccessTokenRequest{
			Code:  "test",
			State: "",
		})

		if assert.NoError(t, err) {
			assert.Equal(t, testUnionId, resp.GetIdentity().GetUnionId())
			assert.Equal(t, testOpenId, resp.GetIdentity().GetOpenId())
			assert.Equal(t, testSessionKey, resp.GetToken().GetAccessToken())
		}
	})

	t.Run("Token failed", func(t *testing.T) {
		_, err := cli.AccessToken(ctx, &api.AccessTokenRequest{
			Code:  "error",
			State: "",
		})

		assert.Error(t, err)
	})
}
