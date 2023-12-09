package main

import (
	"net/http/httputil"
	"net/url"

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
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   "localhost:4200",
		})
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})

	// This handler will return the default application
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"message": "Not Found!",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
