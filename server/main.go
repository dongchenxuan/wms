package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
project-root/
  |- main.go
  |- dist/
      |- static
      	  |- css
      	  |- fonts
      	  |- img
      	  |- js
      	      |- vendor.5afa854d477e17faad12.js
      |- index.html
*/

//go:embed dist/*
var f embed.FS

func serveIndex(c *gin.Context) {
	data, err := f.ReadFile("dist/index.html")
	if err != nil {
		c.String(http.StatusNotFound, "File not found: %v", err)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}

func main() {
	router := gin.Default()

	// Serve static files from the embedded dist/static directory
	router.StaticFS("/static", http.Dir("./dist/static"))

	// Handle root path and /foo path by serving index.html from the embedded directory
	router.GET("/", serveIndex)
	router.GET("/foo", serveIndex)

	router.Run(":9080")
}
