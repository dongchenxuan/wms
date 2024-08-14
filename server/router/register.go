package router

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"wms/code"
)

// Handler handler接口，提供一个注册函数
type Handler interface {
	RegisterFunc(engine *gin.Engine)
}

// JSONResp 给前端的json格式返回
type JSONResp struct {
	Timestamp int64       `json:"timestamp"`
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
}

type JSONErrorResp struct {
	Timestamp int64  `json:"timestamp"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
}

// AlphaHandlerFunc todo:返回json结果的处理函数
type AlphaHandlerFunc func(*gin.Context) (int, interface{}, error)

// JSONWrapper todo: 统一接口函数
func JSONWrapper(fn AlphaHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var aCode int
		var data interface{}
		var err error
		aCode, data, err = fn(c)
		if err == nil && aCode == code.HttpOk {
			render(c, http.StatusOK, JSONResp{
				Timestamp: time.Now().Unix(),
				Code:      aCode,
				Data:      data,
			})
		} else {
			render(c, http.StatusOK, JSONErrorResp{
				Timestamp: time.Now().Unix(),
				Code:      aCode,
				Msg:       err.Error(),
			})
		}
	}
}

// render todo: 输出
func render(c *gin.Context, code int, data interface{}) {
	if encoding := c.GetHeader("Content-Encoding"); encoding == "gzip" {
		bytes, err := json.Marshal(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		gzip, err := Gzip(bytes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.Data(code, "application/gzip", gzip)
	} else {
		c.JSON(code, data)
	}
}

// Gzip todo: gzip压缩，压缩算法为Default，传入待压缩字节流，返回压缩后字节流或错误
func Gzip(b []byte) ([]byte, error) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	_, err := gw.Write(b)
	if err != nil {
		logrus.Error("compress error=%v", err)
		return nil, err
	}
	if err = gw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
