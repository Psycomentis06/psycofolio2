package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/psycomentis/psycofolio++/src/config"
	"gorm.io/gorm"
)

func Injector(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		config.InjectGormDB(ctx, db)
		ctx.Next()
	}
}
