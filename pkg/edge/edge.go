package edge

import (
	"errors"
	"net/http"
	"strings"

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
		// Login page
		acGroup.GET(loginHtml, func(c *gin.Context) {
			ctx := c.Request.Context()
			info, err := e.AccountSvc.Info(ctx, &empty.Empty{})
			if err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
			}
			c.HTML(http.StatusOK, loginHtml, gin.H{
				"basePath": e.Info.BasePath,
				"name":     e.Info.Name,
				"desc":     e.Info.Desc,
				"info":     info,
			})
		})
		// Info api
		acGroup.GET("", func(c *gin.Context) {
			ctx := c.Request.Context()
			info, err := e.AccountSvc.Info(ctx, &empty.Empty{})
			if err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"basePath": e.Info.BasePath,
				"name":     e.Info.Name,
				"desc":     e.Info.Desc,
				"info":     info,
			})
		})

		// Logout
		acGroup.GET("logout", func(c *gin.Context) {
			token, err := checkToken(c)
			if err != nil {
				_ = c.AbortWithError(http.StatusUnauthorized, err)
				return
			}

			ctx := c.Request.Context()
			resp, err := e.AccountSvc.Logout(ctx, &api.LogoutRequest{
				Token: token,
			})
			handleJsonResp(c, err, resp)
		})

		// Refresh token
		acGroup.POST("refresh", func(c *gin.Context) {
			token, err := checkToken(c)
			if err != nil {
				_ = c.AbortWithError(http.StatusUnauthorized, err)
				return
			}

			ctx := c.Request.Context()
			newToken, err := e.AccountSvc.Refresh(ctx, &api.RefreshRequest{
				Token: token,
			})
			handleJsonResp(c, err, newToken)
		})

		// Account auth client api
		acCliGroup := acGroup.Group("/client/:clientId")
		acCliGroup.GET("authorize", func(c *gin.Context) {
			clientId := c.Param("clientId")
			var req struct {
				State        string `form:"state"`
				RedirectUrl  string `form:"redirectUrl"`
				Redirect     bool   `form:"redirect"`
				ResponseType string `form:"responseType"`
				Scope        string `form:"scope"`
			}
			if err := c.BindQuery(&req); err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			ctx := c.Request.Context()
			resp, err := e.AccountSvc.Login(ctx, &api.LoginRequest{
				ResponseType: req.ResponseType,
				ProviderKey:  clientId,
				RedirectUrl:  req.RedirectUrl,
				Scope:        req.Scope,
				State:        req.State,
			})

			if err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
			}

			if req.Redirect {
				c.Redirect(http.StatusSeeOther, resp.GetRedirectTo())
			}

			handleJsonResp(c, err, resp)
		})

		acCliGroup.GET("callback", func(c *gin.Context) {
			clientId := c.Param("clientId")

			ctx := c.Request.Context()
			resp, err := e.AccountSvc.Callback(ctx, &api.CallbackRequest{
				State:       c.Query("state"),
				Code:        c.Query("code"),
				ProviderKey: clientId,
			})

			handleJsonResp(c, err, resp)
		})
	}

	// 业务api
}

func handleJsonResp(c *gin.Context, err error, resp interface{}) {
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, Success(resp))
}

func checkToken(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")
	if strings.TrimSpace(token) == "" {
		return "", errors.New("Not authorize ")
	}

	if strings.HasPrefix(token, "Bearer ") {
		token = strings.Replace(token, "Bearer ", "", -1)
	}
	return token, nil
}

func (e *Edge) Namespace() string {
	return e.Info.BasePath
}

func NewEdge() Api {
	return &Edge{}
}
