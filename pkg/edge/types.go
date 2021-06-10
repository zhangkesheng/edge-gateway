package edge

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Api interface {
	Router(r gin.IRouter)
	Namespace() string
}

type ResultWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *ResultWrapper {
	return &ResultWrapper{
		Code:    200,
		Message: "OK",
		Data:    data,
	}
}
