package oauth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"github.com/zhangkesheng/edge-gateway/api/v1"
	"github.com/zhangkesheng/edge-gateway/pkg/utils"
)

type DingTalkLoginService struct {
	config Config
}

func (d *DingTalkLoginService) Auth(ctx context.Context, req *api.AuthRequest) (*api.AuthResponse, error) {
	params := url.Values{
		"response_type": {"code"},
		"appid":         {d.config.ClientId},
		"scope":         {"snsapi_login"},
		"state":         {req.GetState()},
	}

	if len(strings.TrimSpace(req.GetRedirectUrl())) > 0 {
		params.Add("redirect_uri", req.GetRedirectUrl())
	} else {
		params.Add("redirect_uri", d.config.DefaultRedirect)
	}

	return &api.AuthResponse{
		RedirectTo: fmt.Sprintf("%s?%s", d.config.AuthUrl, params.Encode()),
	}, nil
}

func (d *DingTalkLoginService) AccessToken(ctx context.Context, req *api.AccessTokenRequest) (*api.AccessTokenResponse, error) {
	onError := func(err error) (*api.AccessTokenResponse, error) {
		return nil, errors.Wrap(err, "DingTalkLoginService Login")
	}

	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	signature := utils.HmacSha256Sign(d.config.Secret, timestamp)

	params := url.Values{}
	params.Add("signature", signature)
	params.Add("timestamp", timestamp)
	params.Add("accessKey", d.config.ClientId)

	reqBody := map[string]string{
		"tmp_auth_code": req.GetCode(),
	}

	b, err := json.Marshal(&reqBody)
	if err != nil {
		return onError(err)
	}

	tokenReq, err := http.NewRequest("POST", fmt.Sprintf("%s?%s", d.config.AccessTokenUrl, params.Encode()), bytes.NewReader(b))
	if err != nil {
		return onError(err)
	}
	tokenReq.Header.Set("Content-Type", "application/json")

	return doAuthRequest(tokenReq, func(result gjson.Result) (*api.AccessTokenResponse, error) {
		if result.Get("errcode").Int() != 0 {
			return nil, errors.New(result.Get("errmsg").String())
		}
		userInfo := result.Get("user_info")
		return &api.AccessTokenResponse{
			Identity: &api.Identity{
				OpenId:  userInfo.Get("openid").String(),
				UnionId: userInfo.Get("unionid").String(),
				Nick:    userInfo.Get("nick").String(),
				Source:  "dingtalk",
				Avatar:  "",
				Email:   "",
			},
			Raw: userInfo.String(),
		}, nil
	})
}

func (d *DingTalkLoginService) RefreshToken(ctx context.Context, req *api.RefreshTokenRequest) (*api.RefreshTokenResponse, error) {
	return nil, nil
}

func (d *DingTalkLoginService) Profile(ctx context.Context, req *api.ProfileRequest) (*api.ProfileResponse, error) {
	return nil, nil
}

func NewDingTalk(config Config) api.OAuthClientServer {
	return &DingTalkLoginService{config: config}
}
