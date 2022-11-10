package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_project_demo/common"
)

// Response json响应中间件
func Response(c *gin.Context) {
	common.SetResponseContext(c, common.NewOKResponse(), nil)
	c.Next()
	rsp, _ := common.GetResponseContext(c)
	c.JSON(http.StatusOK, rsp)
}
