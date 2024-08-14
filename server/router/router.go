package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wms/config"
	"wms/log"
	"wms/pages"
)

var (
	handlers []Handler
)

func Init() {
	var engine *gin.Engine

	if config.Instance.Env == config.Delos {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
	} else {
		gin.SetMode(gin.DebugMode)
		engine = gin.Default()
	}

	// Middleware
	engine.Use(NewParamsMiddleware())

	// Serve static files from the web directory
	engine.StaticFS("/static", http.Dir("../web/dist/static"))

	// Web page
	registerPages(engine)

	// Api
	api := engine.Group("/api/v1/wms")
	api.GET("/version", JSONWrapper(version))
	initHandler()
	registerHandlers(engine)

	err := engine.Run(config.Instance.Http.Listen)
	if err != nil {
		log.Errorf("[gin] init error:%v", err)
		panic(err)
	}
}

// todo: loading web pages
func registerPages(engine *gin.Engine) {
	engine.GET("/", pages.Index)
}

// todo: init handler
func initHandler() {
	//handlers = append(handlers, NewHandler())
}

// todo: register handlers
func registerHandlers(engine *gin.Engine) {
	for _, h := range handlers {
		h.RegisterFunc(engine)
	}
}
