package common

import "github.com/gin-gonic/gin"

const (
	HeaderKeyForAuthToken = "AUTH_TOKEN"
)

func GetAuthTokenHeader(c *gin.Context) (token string) {
	token = c.Request.Header.Get(HeaderKeyForAuthToken)
	return
}
