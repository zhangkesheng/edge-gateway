package edge

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zhangkesheng/edge-gateway/pkg/types"
)

type Edge struct {
	Name     string
	Desc     string
	Version  string
	BasePath string

	// 账户体系
	AccountSvc types.ApiRoute

	// 后端服务
	BackendSvc map[string]types.ApiRoute
}

type Backend struct {
	Host string
}

func (edge *Edge) Router(r gin.IRouter) error {
	onError := func(err error) error {
		return errors.Wrap(err, "Edge router err.")
	}

	r.GET("status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// 账户体系
	if edge.AccountSvc != nil {
		if err := edge.AccountSvc.Router(r.Group(edge.AccountSvc.Namespace())); err != nil {
			return onError(err)
		}
	}
	// 业务api
	// TODO:
	// 1. 参数校验
	// 2. 结果过滤
	// ...
	// 先直接透传
	if edge.BackendSvc != nil {
		for bastPath, route := range edge.BackendSvc {
			if err := route.Router(r.Group(bastPath)); err != nil {
				return onError(err)
			}
		}
	}
	return nil
}

func (edge *Edge) Namespace() string {
	return edge.BasePath
}
