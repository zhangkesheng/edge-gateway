package oauth

import (
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

func New(source Source, config Config) api.OAuthClientServer {
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
