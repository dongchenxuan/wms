package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wms/web"
)

// Index todo: 主页
func Index(c *gin.Context) {
	data, err := web.Dist.ReadFile("dist/index.html")
	if err != nil {
		c.String(http.StatusNotFound, "File not found: %v", err)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}
