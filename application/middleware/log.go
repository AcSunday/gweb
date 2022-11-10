package middleware

import (
	"bytes"
	"go_project_demo/component"
	"go_project_demo/library/log"
	"io/ioutil"
	"time"

	"go_project_demo/common"

	"github.com/gin-gonic/gin"
)

// Log 日志中间件
func Log(c *gin.Context) {
	start := time.Now()
	body, err := c.GetRawData()
	if err == nil {
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}

	c.Next()

	end := time.Now()
	latency := end.Sub(start).Microseconds()
	rsp, err := common.GetResponseContext(c)

	fields := log.F{
		"log_type":   common.LogTypeForHTTPRequest,
		"client_ip":  c.ClientIP(),
		"method":     c.Request.Method,
		"uri":        c.Request.RequestURI,
		"body":       string(body),
		"response":   rsp,
		"start":      start,
		"end":        end,
		"latency_ms": latency,
	}

	component.InfLogger.Info(fields)

	if err != nil {
		fields["error"] = err
		component.ErrLogger.Error(fields)
	}
}
