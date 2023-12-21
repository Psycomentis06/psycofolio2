package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/psycomentis/psycofolio++/src/middlewares"
	"github.com/psycomentis/psycofolio++/src/services"
	"github.com/rs/zerolog/log"
)

func main() {
	services.InitApplicationFolders()
	cnf, err := services.LoadConfig(services.ConfigFile)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load configuration")
		return
	}
	db, dbErr := services.CreateDBInstance(&cnf)
	if dbErr != nil {
		log.Err(dbErr)
		return
	}
	services.Migrate(db)

	r := gin.Default()
	r.Use(middlewares.Injector(db, &cnf))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Psycofolio++",
		})
	})

	// This handler will return the admin Angular app
	r.GET("/admin/*any", func(ctx *gin.Context) {
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

	r.Run("0.0.0.0:" + cnf.ServerPort)
}
