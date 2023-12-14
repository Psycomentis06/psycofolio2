package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// This handler will return the admin Angular app
	r.GET("/admin/*any", func(ctx *gin.Context) {
		/*proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   "localhost:4200",
		})
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
		*/
		r.LoadHTMLGlob("web/admin/dist/admin/browser/*")
		ctx.HTML(200, "index.html", nil)
	})

	// This handler will return the default application
	r.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if strings.Contains(path, ".css") || strings.Contains(path, ".js") {
			path = strings.Replace(path, "/", "", 1)
			ctx.File(fmt.Sprintf("web/admin/dist/admin/browser/%s", path))
			return
		}
		ctx.JSON(404, gin.H{
			"message": "Not Found!",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
