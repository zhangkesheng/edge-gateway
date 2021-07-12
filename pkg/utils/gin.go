package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zhangkesheng/edge-gateway/pkg/types"
)

func CheckToken(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")
	if strings.TrimSpace(token) == "" {
		return "", errors.New("Not authorize ")
	}

	if strings.HasPrefix(token, "Bearer ") {
		token = strings.Replace(token, "Bearer ", "", -1)
	}
	return token, nil
}

func HandleJsonResp(c *gin.Context, err error, resp interface{}) {
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, types.Success(resp))
}
