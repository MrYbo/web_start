package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_start/app/container/user"
	"web_start/app/middleware"
)

type routerGroup func(engine *gin.Engine)

var routers = [...]routerGroup{
	user.Router,
}

func Init(engine *gin.Engine) *gin.Engine {
	// 配置模板
	engine.LoadHTMLGlob("app/view/*")

	engine.NoRoute(middleware.Auth(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 路由初始化
	for _, router := range routers {
		router(engine)
	}
	return engine
}
