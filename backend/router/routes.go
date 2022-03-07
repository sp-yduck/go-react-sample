package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sp-yduck/go-react-sample/backend/api"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("build/index.html")
	r.Static("/static", "build/static")
	r.StaticFile("/favicon.ico", "build/favicon.ico")
	r.StaticFile("/logo192.png", "build/logo192.png")
	r.StaticFile("/logo512.png", "build/logo512.png")
	r.StaticFile("/manifest.json", "build/manifest.json")

	r.StaticFile("/", "build/index.html")
	r.StaticFile("/hello", "build/index.html")
	r.StaticFile("/world", "build/index.html")

	r.GET("/api", api.Hello)

	return r
}
