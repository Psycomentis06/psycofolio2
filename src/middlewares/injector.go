package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/psycomentis/psycofolio++/src/services"
	"gorm.io/gorm"
)

func Injector(db *gorm.DB, cnf *services.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		services.InjectGormDB(ctx, db)
		services.InjectConfig(ctx, cnf)
		ctx.Next()
	}
}
