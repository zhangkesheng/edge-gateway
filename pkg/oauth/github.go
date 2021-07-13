package oauth

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

type GithubService struct {
	config config
}

func (g *GithubService) Auth(ctx context.Context, req *api.AuthRequest) (*api.AuthResponse, error) {
	params := url.Values{
		"response_type": {"code"},
		"client_id":     {g.config.clientId},
		"state":         {req.GetState()},
	}

	if len(strings.TrimSpace(req.GetScope())) > 0 {
		params.Add("scope", req.GetScope())
	} else {
		params.Add("scope", g.config.defaultScope)
	}

	if len(strings.TrimSpace(req.GetRedirectUrl())) > 0 {
		params.Add("redirect_uri", req.GetRedirectUrl())
	} else {
		params.Add("redirect_uri", g.config.defaultRedirect)
	}

	return &api.AuthResponse{
		RedirectTo: fmt.Sprintf("%s?%s", g.config.authUrl, params.Encode()),
	}, nil
}

func (g *GithubService) AccessToken(ctx context.Context, req *api.AccessTokenRequest) (*api.AccessTokenResponse, error) {
	onError := func(err error) (*api.AccessTokenResponse, error) {
		return nil, errors.Wrap(err, "GithubService.Login")
	}

	params := url.Values{
		"code":          []string{req.GetCode()},
		"client_id":     []string{g.config.clientId},
		"client_secret": []string{g.config.secret},
		"grant_type":    []string{"authorization_code"},
	}

	tokenReq, err := http.NewRequest("POST", fmt.Sprintf("%s?%s", g.config.accessTokenUrl, params.Encode()), nil)
	if err != nil {
		return onError(err)
	}
	tokenReq.Header.Set("Content-Type", "application/json")
	tokenReq.Header.Set("Accept", "application/json")

	return doAuthRequest(tokenReq, func(result gjson.Result) (*api.AccessTokenResponse, error) {
		if len(result.Get("error").String()) != 0 {
			return nil, errors.New(result.Get("error_description").String())
		}

		return &api.AccessTokenResponse{
			Token: &api.Token{
				AccessToken: result.Get("access_token").String(),
				Scope:       result.Get("scope").String(),
				TokenType:   result.Get("token_type").String(),
			},
			Raw: result.String(),
		}, nil
	})
}

func (g *GithubService) RefreshToken(ctx context.Context, req *api.RefreshTokenRequest) (*api.RefreshTokenResponse, error) {
	return nil, nil
}

func (g *GithubService) Profile(ctx context.Context, req *api.ProfileRequest) (*api.ProfileResponse, error) {
	onError := func(err error) (*api.ProfileResponse, error) {
		return nil, errors.Wrap(err, "GithubService.Profile")
	}

	tokenReq, err := http.NewRequest("GET", g.config.apiUrl, nil)
	if err != nil {
		return onError(err)
	}
	tokenReq.Header.Set("Content-Type", "application/json")
	tokenReq.Header.Set("Authorization", "token "+req.GetAccessToken())

	resp, err := http.DefaultClient.Do(tokenReq)
	if err != nil {
		return onError(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return onError(err)
	}

	result := gjson.ParseBytes(body)

	errMsg := result.Get("message").String()
	if len(errMsg) != 0 {
		return onError(errors.New(errMsg))
	}

	return &api.ProfileResponse{
		Raw: result.String(),
		Identity: &api.Identity{
			OpenId:  result.Get("id").String(),
			UnionId: result.Get("id").String(),
			Nick:    result.Get("name").String(),
			Source:  "github",
			Avatar:  result.Get("avatar_url").String(),
			Email:   result.Get("email").String(),
		},
	}, nil
}

func NewGithub(config config) api.OAuthClientServer {
	return &GithubService{config: config}
}
