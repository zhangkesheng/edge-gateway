package backend

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zhangkesheng/edge-gateway/pkg/types"
)

type ReverseProxy struct {
	Rewrite map[string]string
	Host    string
	// Support all methods. Include: GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT
	// TODO support all path: `*`
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
			path := c.Request.URL.String()
			for k, v := range proxy.Rewrite {
				path = strings.Replace(path, fmt.Sprintf("/%s", k), v, 1)
			}
			c.Request.URL, _ = url.Parse(path)
			p.ServeHTTP(c.Writer, c.Request)
		})
	}

	return nil
}

func (proxy *ReverseProxy) Namespace() string {
	return ""
}

func NewReverseProxy(host string, apis []string, edgeBasePath string, rewrite map[string]string) types.ApiRoute {
	if _, ok := rewrite[edgeBasePath]; !ok {
		rewrite[edgeBasePath] = ""
	}
	return &ReverseProxy{
		Rewrite: rewrite,
		Host:    host,
		Apis:    apis,
	}
}
