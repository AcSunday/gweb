package middleware

import (
	"github.com/gin-gonic/gin"
	"go_project_demo/common"
	"go_project_demo/component"
	"go_project_demo/library/log"
	"net/http"
)

func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			component.ErrLogger.Error(log.F{
				"log_type": common.LogTypeForPanic,
				"info":     err,
			})
			c.Status(http.StatusInternalServerError)
		}
	}()

	c.Next()
}
