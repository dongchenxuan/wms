package router

import (
	"github.com/gin-gonic/gin"
	"wms/code"
)

func version(ctx *gin.Context) (int, interface{}, error) {
	Data := map[string]string{
		"name":    "wms",
		"version": "1.0.0.0",
	}

	return code.HttpOk, Data, nil
}
