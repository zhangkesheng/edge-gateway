package edge

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhangkesheng/edge-gateway/pkg/app"
)

type Edge struct {
	Name     string
	Desc     string
	Version  string
	BasePath string

	AccountSvc app.Api
}

func (edge *Edge) Router(r gin.IRouter) {
	r.GET("status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// 账户体系
	if edge.AccountSvc != nil {
		edge.AccountSvc.Router(r.Group(edge.AccountSvc.Namespace()))
	}
	// 业务api
}

func (edge *Edge) Namespace() string {
	return edge.BasePath
}
