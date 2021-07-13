package types

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type ApiRoute interface {
	Router(r gin.IRouter) error
	Namespace() string
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

func InternalErr(err error) *ResultWrapper {
	return &ResultWrapper{
		Code:    http.StatusInternalServerError,
		Message: "ERROR",
		Data:    err.Error(),
	}
}
