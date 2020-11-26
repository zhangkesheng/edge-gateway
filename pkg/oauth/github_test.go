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
	testGithubToken = "e72e16c7e42f292c6912e7710c838347ae178b4a"
)

func TestGithub(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login/oauth/access_token" {
			code := r.URL.Query().Get("code")
			if code == "test" {
				_, _ = w.Write([]byte(fmt.Sprintf(`{"access_token":"%s", "scope":"repo,gist", "token_type":"bearer"}`, testGithubToken)))
			} else {
				_, _ = w.Write([]byte(`{"error":"bad_verification_code","error_description":"The code passed is incorrect or expired.","error_uri":"https://docs.github.com/apps/managing-oauth-apps/troubleshooting-oauth-app-access-token-request-errors/#bad-verification-code"}`))
			}
		} else {
			_, _ = w.Write([]byte(`{"login":"test","id":123456,"node_id":"tttt","type":"User","site_admin":false,"name":"Test","company":null,"blog":"","location":null,"email":null,"hireable":null,"twitter_username":null}`))
		}
	}))

	defer ts.Close()

	config := Config{
		ClientId:        "test",
		Secret:          "test",
		AuthUrl:         "https://github.com/login/oauth/authorize",
		LogoutUrl:       "",
		AccessTokenUrl:  ts.URL + "/login/oauth/access_token",
		ApiUrl:          ts.URL + "/user",
		DefaultRedirect: "https://www.bestzks.com",
		DefaultScope:    "",
	}
	cli := New(GitHub, config)

	ctx := context.Background()
	t.Run("Auth url", func(t *testing.T) {
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
			assert.Equal(t, testGithubToken, resp.GetToken().GetAccessToken())
		}
	})

	t.Run("Token failed", func(t *testing.T) {
		_, err := cli.AccessToken(ctx, &api.AccessTokenRequest{
			Code:  "error",
			State: "",
		})

		assert.Error(t, err)
	})

	t.Run("User profile", func(t *testing.T) {
		resp, err := cli.Profile(ctx, &api.ProfileRequest{
			AccessToken: "ttttttttttttttttt",
		})

		if assert.NoError(t, err) {
			assert.Equal(t, "Test", resp.GetIdentity().GetNick())
			assert.Equal(t, "123456", resp.GetIdentity().GetOpenId())
		}
	})
}
