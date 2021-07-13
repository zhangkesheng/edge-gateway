package types

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

type ApiRoute interface {
	Router(r gin.IRouter) error
	Namespace() string
}

type AccountRouter interface {
	ApiRoute
	api.AccountServer
}

type ResultWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *ResultWrapper {
	return &ResultWrapper{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    data,
	}
}

func CtxError(c *gin.Context, status int, err error) {
	c.AbortWithStatusJSON(status, Error(status, err))
}

func Error(status int, err error) *ResultWrapper {
	return &ResultWrapper{
		Code:    status,
		Message: "ERROR",
		Data:    err.Error(),
	}
}

func InternalErr(err error) *ResultWrapper {
	return Error(http.StatusInternalServerError, err)
}
