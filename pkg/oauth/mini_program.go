package oauth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

type MiniProgramService struct {
	config config
}

type code2SessionResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func (srv *MiniProgramService) Auth(ctx context.Context, req *api.AuthRequest) (*api.AuthResponse, error) {
	return &api.AuthResponse{}, nil
}

func (srv *MiniProgramService) AccessToken(ctx context.Context, req *api.AccessTokenRequest) (*api.AccessTokenResponse, error) {
	onError := func(err error) (*api.AccessTokenResponse, error) {
		return nil, errors.Wrap(err, "Code2Session")
	}

	tokenReq, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%s/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
			srv.config.accessTokenUrl,
			srv.config.clientId,
			srv.config.secret,
			req.GetCode(),
		),
		nil,
	)
	if err != nil {
		return onError(err)
	}
	return doAuthRequest(tokenReq, func(result gjson.Result) (*api.AccessTokenResponse, error) {
		if result.Get("errcode").Int() != 0 {
			return nil, errors.New(result.Get("error_description").String())
		}

		return &api.AccessTokenResponse{
			Token: &api.Token{
				AccessToken: result.Get("session_key").String(),
			},
			Raw: result.String(),
			Identity: &api.Identity{
				OpenId:  result.Get("openid").String(),
				UnionId: result.Get("unionid").String(),
				Source:  "mini-program",
				Avatar:  "",
				Email:   "",
			},
		}, nil
	})
}

func (srv *MiniProgramService) RefreshToken(ctx context.Context, req *api.RefreshTokenRequest) (*api.RefreshTokenResponse, error) {
	return &api.RefreshTokenResponse{}, nil
}

func (srv *MiniProgramService) Profile(ctx context.Context, req *api.ProfileRequest) (*api.ProfileResponse, error) {
	return &api.ProfileResponse{}, nil
}

func NewMiniProgram(config config) api.OAuthClientServer {
	return &MiniProgramService{config: config}
}
