package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sp-yduck/go-react-sample/backend/api"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("../flontend/build/index.html")
	r.Static("/static", "../flontend/build/static")
	r.StaticFile("/favicon.ico", "../flontend/build/favicon.ico")
	r.StaticFile("/logo192.png", "../flontend/build/logo192.png")
	r.StaticFile("/logo512.png", "../flontend/build/logo512.png")
	r.StaticFile("/manifest.json", "../flontend/build/manifest.json")

	r.StaticFile("/", "../flontend/build/index.html")
	r.StaticFile("/hello", "../flontend/build/index.html")
	r.StaticFile("/world", "../flontend/build/index.html")

	r.GET("/api", api.Hello)

	return r
}
