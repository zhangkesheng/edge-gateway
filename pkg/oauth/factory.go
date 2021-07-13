package oauth

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

func New(source Source, config config) api.OAuthClientServer {
	switch source {
	case DingTalkLogin:
		config.authUrl = "https://oapi.dingtalk.com/connect/qrconnect"
		config.accessTokenUrl = "https://oapi.dingtalk.com/sns/getuserinfo_bycode"
		return NewDingTalk(config)
	case GitHub:
		config.authUrl = "https://github.com/login/oauth/authorize"
		config.accessTokenUrl = "https://github.com/login/oauth/access_token"
		config.apiUrl = "https://api.github.com/user"
		return NewGithub(config)
	case MiniProgram:
		return NewMiniProgram(config)
	default:
		return nil
	}
}

func NewOauth(option Option) api.OAuthClientServer {
	return New(option.Source, config{
		clientId:        option.ClientId,
		secret:          option.Secret,
		defaultRedirect: option.DefaultRedirect,
	})
}

type Option struct {
	Source          Source
	ClientId        string
	Secret          string
	DefaultRedirect string
}

type Source int

const (
	UnKnown Source = iota

	// 需要自己实现的部分
	MobileVerify
	UserNamePwd

	// 第三方服务
	DingTalkLogin
	GitHub
	MiniProgram
)

type config struct {
	clientId        string
	secret          string
	authUrl         string
	logoutUrl       string
	accessTokenUrl  string
	apiUrl          string
	defaultRedirect string
	defaultScope    string
}

func doAuthRequest(req *http.Request, handler func(result gjson.Result) (*api.AccessTokenResponse, error)) (*api.AccessTokenResponse, error) {
	onError := func(err error) (*api.AccessTokenResponse, error) {
		return nil, errors.Wrap(err, "Auth.Request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return onError(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return onError(err)
	}

	result := gjson.ParseBytes(body)

	if resp, err := handler(result); err != nil {
		return onError(err)
	} else {
		return resp, nil
	}
}
