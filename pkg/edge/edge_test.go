package edge

import (
	"database/sql"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"github.com/zhangkesheng/edge-gateway/pkg/account"
)

func TestEdge(t *testing.T) {
	mr, _ := miniredis.Run()

	edge := &Edge{
		Info: Info{
			Name:     "Test Edge",
			Desc:     "A test edge demo.",
			Version:  "v0.0.1",
			BasePath: "test",
		},
		AccountSvc: account.NewAccount(
			"http;//127.0.0.1:8080",
			"Test",
			"Test",
			600,
			redis.NewClient(&redis.Options{
				Addr: mr.Host(),
			}),
			&sql.DB{}),
	}

	router := gin.New()

	edge.Router(router)

	ts := httptest.NewServer(router)
	defer ts.Close()

	doTestRequest := func(url string) string {
		res, err := ts.Client().Get(ts.URL + url)
		if assert.NoError(t, err) {
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			return string(body)
		}
		return ""
	}

	t.Run("status", func(t *testing.T) {
		resp := doTestRequest("/status")
		assert.Equal(t, "OK", gjson.Get(resp, "status").String())
	})
}
