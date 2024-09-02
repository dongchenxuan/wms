package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Icon todo: Icon
func Icon(c *gin.Context) {
	data, err := Dist.ReadFile("dist/favicon.ico")
	if err != nil {
		c.String(http.StatusNotFound, "File not found: %v", err)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}

// Index todo: 主页
func Index(c *gin.Context) {
	data, err := Dist.ReadFile("dist/index.html")
	if err != nil {
		c.String(http.StatusNotFound, "File not found: %v", err)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}

// NotFound todo: 404
func NotFound(c *gin.Context) {

}
