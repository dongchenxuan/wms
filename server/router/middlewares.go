package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wms/code"
)

// NewParamsMiddleware TODO: 解析http请求参数的中间件
func NewParamsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			render(c, http.StatusOK, JSONErrorResp{
				Timestamp: time.Now().Unix(),
				Code:      code.ServerErr,
				Msg:       code.ErrMsgMap[code.ServerErr].Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
