package router

import (
	"github.com/gin-gonic/gin"
	"io/fs"
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
	fe, err := fs.Sub(pages.Dist, "dist/static")
	if err != nil {
		panic(err)
	}
	engine.StaticFS("/static", http.FS(fe))

	// Web page
	registerPages(engine)

	// Api
	api := engine.Group("/api/v1/wms")
	api.GET("/version", JSONWrapper(version))
	initHandler()
	registerHandlers(engine)

	err = engine.Run(config.Instance.Http.Listen)
	if err != nil {
		log.Errorf("[gin] init error:%v", err)
		panic(err)
	}
}

// todo: loading web pages
func registerPages(engine *gin.Engine) {
	engine.GET("/", pages.Index)
	engine.GET("/favicon.ico", pages.Icon)
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
