package controller

import (
	"github.com/gin-gonic/gin"
	"go_project_demo/application/logic"
	"go_project_demo/common"
)

// SayHelloRequest 请求示例
type SayHelloRequest struct {
	FirstName string `json:"first_name" form:"first_name" binding:"required,lte=128"`
	LastName  string `json:"last_name" form:"last_name" binding:"lte=128"`
}

func SayHello(c *gin.Context) {
	req := &SayHelloRequest{}
	if err := c.ShouldBind(req); err != nil {
		common.SetResponseContext(c, common.NewParamErrResponse(err), nil)
		return
	}
	rsp, err := logic.SayHello(req.FirstName, req.LastName)
	common.SetResponseContext(c, rsp, err)
	return
}
