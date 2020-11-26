package oauth

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

const (
	testNick    = "张三"
	testOpenId  = "liSii8KCxxxxx"
	testUnionId = "7Huu46kk"
)

func TestDingTalk(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		code := gjson.ParseBytes(body).Get("tmp_auth_code").String()
		if code == "test" {
			_, _ = w.Write([]byte(fmt.Sprintf(`{"errcode":0,"errmsg":"ok","user_info":{"nick":"%s","openid":"%s","unionid":"%s"}}`, testNick, testOpenId, testUnionId)))
		} else {
			_, _ = w.Write([]byte(`{"errcode":111,"errmsg":"error"}`))
		}
	}))

	defer ts.Close()

	config := Config{
		ClientId:        "test",
		Secret:          "test",
		AuthUrl:         "https://oapi.dingtalk.com/connect/qrconnect",
		LogoutUrl:       "",
		AccessTokenUrl:  ts.URL + "/sns/getuserinfo_bycode",
		ApiUrl:          "",
		DefaultRedirect: "https://www.bestzks.com",
		DefaultScope:    "",
	}
	cli := New(DingTalkLogin, config)

	ctx := context.Background()
	t.Run("Auth", func(t *testing.T) {
		_, err := cli.Auth(ctx, &api.AuthRequest{
			Scope:        "",
			ResponseType: "code",
			RedirectUrl:  "",
			State:        "xxx",
		})

		assert.NoError(t, err)
	})

	t.Run("Token", func(t *testing.T) {
		resp, err := cli.AccessToken(ctx, &api.AccessTokenRequest{
			Code:  "test",
			State: "",
		})

		if assert.NoError(t, err) {
			assert.Equal(t, testNick, resp.GetIdentity().GetNick())
			assert.Equal(t, testOpenId, resp.GetIdentity().GetOpenId())
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
