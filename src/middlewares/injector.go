package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/psycomentis/psycofolio++/src/services"
	"gorm.io/gorm"
)

func Injector(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		services.InjectGormDB(ctx, db)
		ctx.Next()
	}
}
