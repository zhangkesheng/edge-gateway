package edge

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

const (
	loginHtml = "login.html"
)

type Edge struct {
	Info       Info
	AccountSvc api.AccountServer
}

type Info struct {
	Name     string
	Desc     string
	Version  string
	BasePath string
}

func (e *Edge) Router(r gin.IRouter) {
	r.GET("status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// 账户体系
	if e.AccountSvc != nil {
		acGroup := r.Group("account")
		// Login html
		acGroup.GET(loginHtml, func(c *gin.Context) {
			ctx := c.Request.Context()

			if info, err := e.AccountSvc.Info(ctx, &empty.Empty{}); err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
			} else {
				c.HTML(http.StatusOK, loginHtml, map[string]interface{}{
					"basePath": e.Info.BasePath,
					"name":     e.Info.Name,
					"desc":     e.Info.Desc,
					"info":     info,
				})
			}
		})
	}

	// 业务api
}

func (e *Edge) Namespace() string {
	return e.Info.BasePath
}

func NewEdge() Api {
	return &Edge{}
}
