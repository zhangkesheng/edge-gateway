package backend

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zhangkesheng/edge-gateway/pkg/types"
)

type ReverseProxy struct {
	Host string
	Apis []string
}

func (proxy *ReverseProxy) Router(r gin.IRouter) error {
	onError := func(err error) error {
		return errors.Wrap(err, "Backend router err.")
	}
	host, err := url.Parse(proxy.Host)

	if err != nil {
		return onError(err)
	}
	p := httputil.NewSingleHostReverseProxy(host)

	for _, api := range proxy.Apis {
		r.Any(api, func(c *gin.Context) {
			p.ServeHTTP(c.Writer, c.Request)
		})
	}

	return nil
}

func (proxy *ReverseProxy) Namespace() string {
	return ""
}

func NewReverseProxy(host string, apis []string) types.ApiRoute {
	return &ReverseProxy{
		Host: host,
		Apis: apis,
	}
}
