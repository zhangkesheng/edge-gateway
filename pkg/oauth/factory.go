package oauth

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

// TODO handler function
func (f *Factory) New(source Source, config Config) api.OAuthClientServer {
	switch source {
	case DingTalkLogin:
		return NewDingTalk(config)
	case GitHub:
		return NewGithub(config)
	default:
		return nil
	}
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
)

type Config struct {
	ClientId        string
	Secret          string
	AuthUrl         string
	LogoutUrl       string
	AccessTokenUrl  string
	ApiUrl          string
	DefaultRedirect string
	DefaultScope    string
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
